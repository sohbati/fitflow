package com.caa.config;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.context.annotation.PropertySource;
import org.springframework.core.env.Environment;

/**
 * Created by Reza on 13/09/2018.
 */
@Configuration
@PropertySource("classpath:config.properties")
public class ProjectConfig {

    public static final String PROJECT_CoNFIG_BEAN_NAME = "PConfig";
    @Autowired
    Environment environment;

    @Bean("PConfig")
    public ConfigData getProjectConfig() {
        ConfigData configData = new ConfigData();
        configData.setPictureLocation(environment.getProperty("picture.location"));
        //
        configData.setShrinkedImageWidth(Integer.parseInt(environment.getProperty("shrinked.image.width")));
        configData.setShrinkedImageHeight(Integer.parseInt(environment.getProperty("shrinked.image.height")));
        return  configData;
    }
}
