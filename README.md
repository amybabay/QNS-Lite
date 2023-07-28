# QNS-Lite

This repository is forked from the repository containing the code for QNS-Lite,
a simulator for quantum networks. QNS-Lite was originally developed by Ali
Farahbakhsh and Chen Feng and used for the simulations of the paper
["Opportunistic routing in quantum
networks"][https://ieeexplore.ieee.org/abstract/document/9796816], Ali
Farahbakhsh and Chen Feng. In IEEE INFOCOM 2022-IEEE Conference on Computer
Communications, pp. 490-499, 2022.

## Paper Results

To reproduce the results of the paper, open 'main.go' in the 'experiments'
folder, and specify the experiment that you want to run in the main function
(uncomment the line for that experiment and comment out all others). The
experimental figures of the paper correspond to 'experiment1.go' up to
'experiment4.go', in the same order of appearance in the paper. The data of the
table provided in the paper is gathered in 'experiment6.go'.

To run the experiment, you can run, from inside the `experiments` directory:

```
go run .
```

The output will be saved in `experiments/Data/experimentN.txt`, where `N` is
the number of the experiment you chose to run.

After executing each experiment, you can draw the plot by running the
corresponding Python code in 'experiments/Python scripts'. The plot will be
saved to 'experiments/Python scripts'.

For example, to generate Figure 5 from the paper (corresponding to
experiment1), you can run, from inside the `experiments/Python scripts`
directory:

```
python experiment1.py
```
