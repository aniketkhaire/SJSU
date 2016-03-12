package edu.sjsu.cmpe275.lab1;

public interface TweetStats {

    /**
     * reset all the three measurements.
     */
    void resetStats();

    /**
     * @return the length of longest message attempted  since the beginning or last reset. Can be more than 140.
     * Failed messages are counted for this as well.
     */
    int getLengthOfLongestTweetAttempted();

    /**
     * @return the user who has been followed by the biggest number of different users since the beginning or last reset. If there is a tie,
     * return the 1st of such users based on alphabetical order. If the follow action did not succeed, then
     * it does not count into the stats
     */
    String getMostFollowedUser();

    /**
     * The most productive user is determined by the total length of all the messages successfully tweeted since the beginning
     * or last reset. If there is a tie, return the 1st of such users based on alphabetical order.
     *
     * @return the most productive user.
     */
    String getMostProductiveUser();
}