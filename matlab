clear all;

clc;

ans = pi;

    beta = ans;
    f_beta = sin(beta-0.361) + (sin(0.361)*exp(-beta/0.377));
    d_f_beta = cos(beta-0.361) + (sin(0.367)*(-1/0.377)*exp(-beta/0.377));
    
    ans = beta - (f_beta/d_f_beta);
