//package main.java.edu.sjsu.cmpe275.lab1;
package edu.sjsu.cmpe275.lab1;

//import java.util.HashMap;
import java.util.List;
//import java.util.Map;
import java.util.Set;
import java.util.*;

public class TweetStatsImpl implements TweetStats {

    /***
     * Following is the dummy implementation of methods.
     * Students are expected to complete the actual implementation of these methods as part of lab completion.
     */

    @Override
    public void resetStats() {
        // TODO Auto-generated method stub    	
    	//Clearing all the data
    	RetryAndDoStats.followers.clear();
    	RetryAndDoStats.successfulMessages.clear();
    	RetryAndDoStats.unsuccessfulMessages.clear();
    }
    
    @Override
    public int getLengthOfLongestTweetAttempted() {
        // TODO Auto-generated method stub
    	int length = 0;
    	
    	//finding max length of successful messages
    	for(List<String> data : RetryAndDoStats.successfulMessages.values()){
    		for(String message : data){
    			if(message.length() > length)
    				length = message.length();
    		}
    	}
    	
    	//finding max length of unsuccessful messages
    	for(List<String> data : RetryAndDoStats.unsuccessfulMessages.values()){
    		for(String message : data){
    			if(message.length() > length)
    				length = message.length();
    		}
    	}
    	
        return length;
    }

    @Override
    public String getMostFollowedUser() {
        // TODO Auto-generated method stub
    	int noOfFollowers = 0, followers =0;
    	List<String> userName = new ArrayList<String>();
    	
    	//finding user with maximum number of followers
    	for(String name : RetryAndDoStats.followers.keySet()){ 
    		Set s = new HashSet(RetryAndDoStats.followers.get(name));
    		
    		followers = s.size();
    		
    		//if user with highest followers found, clear the list and add this user
    		if(followers > noOfFollowers){
    			noOfFollowers = followers;
    			userName.clear();
    			
    			//add username to list
    			userName.add(name);    			
    		}
    		//if multiple users exist with highest number of followers
    		else if(followers == noOfFollowers){
    			userName.add(name);
    		}
    		//setting followers to 0 for next user in loop
    		followers = 0;
    	}
    	
    	//if empty, return NULL
    	if(userName.isEmpty())
    		return null;
    	//else, sort and return smallest element
    	Collections.sort(userName);
    	return userName.get(0); 
    }
    
    @Override
    public String getMostProductiveUser() {
        // TODO Auto-generated method stub
    	int length = 0;
    	int maxLength = 0;
    	List<String> userName = new ArrayList<String>();
    	
    	//finding max length of successful messages
    	for(String name : RetryAndDoStats.successfulMessages.keySet()){ 
    		//calculate total length for particular user
    		for(String message : RetryAndDoStats.successfulMessages.get(name)){
    			length += message.length();
    		}
    		if(length > maxLength){
    			maxLength = length;
    			userName.clear();
    			
    			//add username to list
    			userName.add(name);    			
    		}
    		else if(length == maxLength){
    			userName.add(name);
    		}
    		length = 0;
    	}
    	
    	//if empty, return NULL
    	if(userName.isEmpty())
    		return null;
    	//else, sort and return smallest element
    	Collections.sort(userName);
    	return userName.get(0);
    }
}