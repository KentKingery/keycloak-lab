package io.acmecode.client;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

public class App 
{
    private static final Logger logger = LoggerFactory.getLogger(App.class);

    public static void main( String[] args ) 
    {
        logger.info("Client startup");
        logger.info("Client shutdown");
    }

    private static void shutdown() 
    {
        logger.info("Client startup");
    }

    private static void startup() 
    {
        logger.info("Client shutdown");
    }
}
