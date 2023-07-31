package main

import (
	"fmt"

	"example.com/benchmark"
	"example.com/config"
)

func experiment6(bm *benchmark.Benchmarker, itrReqs int, itrSingleReq int, maxItr int) {
	// AB: Set up stats tracking
	averageNOPP := make([]float64, itrReqs)
	varianceNOPP := make([]float64, itrReqs)
	averageOPP := make([]float64, itrReqs)
	varianceOPP := make([]float64, itrReqs)

	// AB: set up p_gen values
	p := [10]float64{0.1, 0.2, 0.3, 0.4, 0.5, 0.6, 0.7, 0.8, 0.9, 1}
	ALWT := make([][]float64, 6)
	for i := 0; i <= 5; i++ {
		ALWT[i] = make([]float64, len(p))
	}

	// AB: Set configuration params
	//algos := [3]string{"modified greedy", "nonoblivious local", "qpass"}
	algos := [2]string{"modified greedy", "qpass"}
	topologies := [1]string{"grid"}
	config.SetPSwap(float64(1))
	config.SetSize(10)
	config.SetLifetime(30)
	config.SetNumRequests(20)

	// AB: In this experiment, we test a range of p_gen values. Iterate over
	// each value we are testing
	for p_genIndex, p_gen := range p {
		config.SetPGen(p_gen)
		fmt.Println("p_gen is", p_gen)
		fmt.Println("config.p_gen is", config.GetConfig().GetPGen())

		for algo := 0; algo < len(algos); algo++ {
			fmt.Println("algorithm is", algos[algo])
			bm.Set(itrSingleReq, algos[algo], topologies[0])
			bm.SetKeepReqs(true)

			// AB: generate itrReqs different sets of requests
			for i := 0; i <= itrReqs-1; i++ {
				//fmt.Println("Average Run:", i)
				bm.RegenerateReqs(itrSingleReq)

				// AB: run the algorithm without opportunism
				config.SetOpportunism(false)
				bm.Start(itrSingleReq, maxItr)
				averageNOPP[i] = AverageWaiting(bm.LinksWaitingTime, maxItr)
				varianceNOPP[i] = VarianceWaiting(bm.LinksWaitingTime, maxItr)
				//fmt.Println(*bm)
				//fmt.Println("NOPP Finished.")

				// AB: run the algorithm with opportunisum
				config.SetOpportunism(true)
				bm.Start(itrSingleReq, maxItr)
				//fmt.Println(*bm)
				averageOPP[i] = AverageWaiting(bm.LinksWaitingTime, maxItr)
				varianceOPP[i] = VarianceWaiting(bm.LinksWaitingTime, maxItr)
			}

			// AB: format is like [alg0_nopp: p0, p1, p2, p3, ...]
			//					  [alg0_opp:  p0, p1, p2, p3, ...]
			//					  [alg1_nopp: p0, p1, p2, p3, ...]
			//                    [alg1_opp:  p0, p1, p2, p3, ...]
			ALWT[2*algo][p_genIndex] = AverageWaiting(averageNOPP, maxItr)
			ALWT[2*algo+1][p_genIndex] = AverageWaiting(averageOPP, maxItr)
			fmt.Println("Average NOPP waiting time is:", AverageWaiting(averageNOPP, maxItr))
			fmt.Println("Average OPP waiting time is:", AverageWaiting(averageOPP, maxItr))
		}
	}

	// AB: write to file
	handleFile(ALWT, "./Data/experiment6.txt")

	fmt.Println(ALWT)
}
