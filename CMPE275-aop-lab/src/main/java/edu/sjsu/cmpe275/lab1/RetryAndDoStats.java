package edu.sjsu.cmpe275.lab1;

import org.aopalliance.intercept.MethodInterceptor;
import org.aopalliance.intercept.MethodInvocation;

public class RetryAndDoStats implements MethodInterceptor {
    /***
     * Following is the dummy implementation of advice.
     * Students are expected to complete the required implementation as part of lab completion.
     */
    public Object invoke(MethodInvocation invocation) throws Throwable {
        System.out.println("Method " + invocation.getMethod() + " is called");
        return invocation.proceed();
    }
}
