package com.caa.report;

import com.caa.model.Person;
import com.caa.model.Program;
import com.caa.util.ImageUtil;
import net.sf.jasperreports.engine.*;
import net.sf.jasperreports.engine.data.JRBeanCollectionDataSource;
import net.sf.jasperreports.engine.design.JasperDesign;
import net.sf.jasperreports.engine.xml.JRXmlLoader;

import javax.imageio.ImageIO;
import java.awt.image.BufferedImage;
import java.io.*;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/**
 * Created by Reza on 20/08/2018.
 */
public class ExportReport {

    /**
     *
     * @param exerciseItems
     * @return BASE64 encoded string
     */
    public static String getProgramExerciseAsImage(List<ProgramExercisesReportDTO> exerciseItems,
                                                   Program program, Person person) {
        ByteArrayOutputStream out = getImageStream(exerciseItems, program, person);
        String imageString = ImageUtil.encodeImage(out);
        return imageString;
    }

    public static byte[] getProgramExerciseAsImageInBytes(
            List<ProgramExercisesReportDTO> exerciseItems,
            Program program, Person person) {
        ByteArrayOutputStream out = getImageStream(exerciseItems, program, person);
        return out.toByteArray();
    }


    private static ByteArrayOutputStream getImageStream(
            List<ProgramExercisesReportDTO> exerciseItems, Program program, Person person) {
        try {
            if (exerciseItems.size() == 0) {
                return null;
            }
            /**
             * Recognize sub-program count
             */
            int subProgramCount = 0;
             if (exerciseItems.get(0).getProgram4Id() != null && exerciseItems.get(0).getProgram4Id()  >0) {
                 subProgramCount = 4;
             }else if (exerciseItems.get(0).getProgram3Id() != null && exerciseItems.get(0).getProgram3Id()  >0) {
                 subProgramCount = 3;
             }else if (exerciseItems.get(0).getProgram2Id() != null && exerciseItems.get(0).getProgram2Id()  >0) {
                 subProgramCount = 2;
             }else {
                 subProgramCount = 1;
             }

            InputStream inputStream = new FileInputStream(
                    new java.io.File(
                           exerciseItems.get(0).getClass().getClassLoader().getResource(
                                   "report/PersonProgramExercises" + subProgramCount + "Session.jrxml").getFile()));

            JasperDesign jasperDesign = JRXmlLoader.load(inputStream);
            JasperReport jasperReport = JasperCompileManager.compileReport(jasperDesign);

            Map<String, Object> parameters = new HashMap();
            parameters.put("personName", person.getFirstName() + ' ' + person.getLastName());
            parameters.put("age", program.getPersonAge());
            parameters.put("tall", program.getPersonTall());
            parameters.put("weight", program.getPersonWeight());

            JasperPrint jasperPrint = JasperFillManager.fillReport(jasperReport, parameters,
                    new JRBeanCollectionDataSource(exerciseItems));

            final String extension = "jpg";
            final float zoom = 2f;

            if (jasperPrint.getPages().size() == 0) {
                return null;
            }

            try(ByteArrayOutputStream out = new ByteArrayOutputStream()){
                BufferedImage image = (BufferedImage) JasperPrintManager.printPageToImage(jasperPrint, 0,zoom);
                ImageIO.write(image, extension, out);

                return out;
            } catch (IOException e) {
                e.printStackTrace();
            }

        } catch (Exception e) {
            e.printStackTrace();
        }
        return null;
    }
}
