package com.caa.util;

import com.caa.config.ConfigData;
import com.caa.config.ProjectConfig;
import org.springframework.context.annotation.AnnotationConfigApplicationContext;

/**
 * Created by Reza on 14/09/2018.
 */
public class PublicUtil {

 private static AnnotationConfigApplicationContext context = null;

 private static AnnotationConfigApplicationContext getContext() {
  if (context == null) {
   context = new AnnotationConfigApplicationContext(ProjectConfig.class);
  }
  return context;
 }

 public static String getProjectConfigFolder() {
  String confFolder = ((ConfigData)getContext().getBean(ProjectConfig.PROJECT_CoNFIG_BEAN_NAME)).getPictureLocation();
  return confFolder;
 }

 public static ConfigData getPublicConfig() {
  return ((ConfigData)getContext().getBean(ProjectConfig.PROJECT_CoNFIG_BEAN_NAME));
 }

}
