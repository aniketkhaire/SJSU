//package main.java.edu.sjsu.cmpe275.lab1;
package edu.sjsu.cmpe275.lab1;

import org.aopalliance.intercept.MethodInterceptor;
import org.aopalliance.intercept.MethodInvocation;
import org.springframework.aop.ProxyMethodInvocation;

import java.io.IOException;
import java.util.*;

public class RetryAndDoStats implements MethodInterceptor {
    /***
     * Following is the dummy implementation of advice.
     * Students are expected to complete the required implementation as part of lab completion.
     */
	
	public static Map<String, List<String>> followers = new HashMap<String, List<String>>();
	public static Map<String, List<String>> successfulMessages = new HashMap<String, List<String>>();
	public static Map<String, List<String>> unsuccessfulMessages = new HashMap<String, List<String>>();	
	
    public Object invoke(MethodInvocation invocation) throws Throwable {
    	
    	String methodName = invocation.getMethod().getName();
    	Object[] arguments = invocation.getArguments();
    	Object obj= null;
    	int count = 0;
    	
    	try{
    		//executing the statement
    		obj = invocation.proceed();
    		
    		//if exception is not thrown, and methodName is "tweet", store the tweet as a successful tweet
    		if(methodName.equals("tweet")){
    			//for "tweet method"
    			String userName = arguments[0].toString();
    			String userTweet = arguments[1].toString();
    			//storing successful tweets
    			storeSuccessfulTweets(userName, userTweet);    			
    		}
    		else{
    			//for "follow method"
    			String follower = arguments[0].toString();
    			String followee = arguments[1].toString();
    			//adding follower to followee
    			addFollower(follower, followee);
    		}
    	}//end of try
    	catch(IOException e){
    		//retry for 3 times
    		while(count++ < 3)
    		{
    			try{
    				invocation.proceed();
    				if(methodName.equals("tweet")){
    					//fetching arguments
    					String userName = arguments[0].toString();
    	    			String userTweet = arguments[1].toString();
    	    			//storing successful tweets
    					storeSuccessfulTweets(userName, userTweet);
    				}
    				else{
    					//fetching arguments
    					String follower = arguments[0].toString();
    	    			String followee = arguments[1].toString();
    	    			//adding follower to followee
    					addFollower(follower, followee);
    				}
    				count = 0;
    				break;
    			}catch(Exception x){
    				if(x instanceof IllegalArgumentException){
    					storeUnsuccessfulTweets(arguments[0].toString(), arguments[1].toString());
    					count = 0;
    					break;
    				}
    				
    				System.out.println("Retrying..."+count);
    			}//end of inner catch
    		}
    		
    		if(count != 0){
    			//unsuccesful message because of IOException
    			if(methodName.equals("tweet"))
        			storeUnsuccessfulTweets(arguments[0].toString(), arguments[1].toString());
    		}

    		//return obj;
    		return invocation.proceed();
    	}//end of outer catch
    	catch(IllegalArgumentException il){
    		//Message length >140, thus adding to unsuccessful tweet
    		storeUnsuccessfulTweets(arguments[0].toString(), arguments[1].toString());
    	}//end of catch
    	
        //System.out.println("Method " + invocation.getMethod() + " is called");
        //return obj;
    	return invocation.proceed();
    }
    
    
    //function to write unsuccessful tweets
    public void storeUnsuccessfulTweets(String userName, String userTweet){
    	/* storing unsuccessful tweets
		 */
    	if(unsuccessfulMessages.containsKey(userName))
		{
			List<String> temp = unsuccessfulMessages.get(userName);
			temp.add(userTweet);
			unsuccessfulMessages.put(userName, temp);
		}else
		{
			List<String> temp = new ArrayList<String>();
			temp.add(userTweet);
			unsuccessfulMessages.put(userName, temp);
		}
    }
    
  //function to write successful tweets
    public void storeSuccessfulTweets(String userName, String userTweet){
    	/* if user already exists, add the tweet to corresponding List<String>
		 * else add the new user to the Map
		 */
		if(successfulMessages.containsKey(userName)){
			List<String> temp = successfulMessages.get(userName);
			temp.add(userTweet);
			successfulMessages.put(userName, temp);
		}else{
			List<String> temp = new ArrayList<String>();
			temp.add(userTweet);
			successfulMessages.put(userName, temp);
		}
    }
    
  //function to add follower to followee
    public void addFollower(String follower, String followee){
    	/* if user already exists, add the follower to corresponding followee
		 * else add the new followee to the Map along with its follower 
		 */
		if(followers.containsKey(followee)){
			List<String> temp = followers.get(followee);
			temp.add(follower);
			followers.put(followee, temp);
		}else{
			List<String> temp = new ArrayList<String>();
			temp.add(follower);
			followers.put(followee, temp);
		}    	
    }    
}