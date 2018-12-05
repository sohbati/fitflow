package com.caa.report;

import com.caa.model.Person;
import com.caa.model.Program;
import com.caa.util.DateUtil;
import com.caa.util.ImageUtil;
import com.caa.util.PublicUtil;
import net.sf.jasperreports.engine.*;
import net.sf.jasperreports.engine.data.JRBeanCollectionDataSource;
import net.sf.jasperreports.engine.design.JasperDesign;
import net.sf.jasperreports.engine.util.JRLoader;
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

    public static byte[] getProgramExerciseAsPDFInBytes(
            List<ProgramExercisesReportDTO> exerciseItems,
            Program program, Person person) {
        try {
            if (exerciseItems.size() == 0) {
                return null;
            }

            JasperPrint jasperPrint = prepareJasperPrint(exerciseItems, program, person);
            if (jasperPrint.getPages().size() == 0) {
                return null;
            }

            byte [] pdfBytes = JasperExportManager.exportReportToPdf(jasperPrint);
            return pdfBytes;
        } catch (Exception e) {
            e.printStackTrace();
        }
        return null;
    }

    private static ByteArrayOutputStream getImageStream(
            List<ProgramExercisesReportDTO> exerciseItems, Program program, Person person) {
        try {
            if (exerciseItems.size() == 0) {
                return null;
            }

            final String extension = "jpg";
            final float zoom = 2f;

            JasperPrint jasperPrint = prepareJasperPrint(exerciseItems, program, person);
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

    private static JasperPrint prepareJasperPrint(
            List<ProgramExercisesReportDTO> exerciseItems, Program program, Person person) throws Exception{
        /**
         * Recognize sub-program count
         */
        int subProgramCount = 0;
        if (exerciseItems.get(0).getProgram6Id() != null && exerciseItems.get(0).getProgram6Id()  >0) {
            subProgramCount = 6;
        }else if (exerciseItems.get(0).getProgram5Id() != null && exerciseItems.get(0).getProgram5Id()  >0) {
            subProgramCount = 5;
        }else if (exerciseItems.get(0).getProgram4Id() != null && exerciseItems.get(0).getProgram4Id()  >0) {
            subProgramCount = 4;
        }else if (exerciseItems.get(0).getProgram3Id() != null && exerciseItems.get(0).getProgram3Id()  >0) {
            subProgramCount = 3;
        }else if (exerciseItems.get(0).getProgram2Id() != null && exerciseItems.get(0).getProgram2Id()  >0) {
            subProgramCount = 2;
        }else {
            subProgramCount = 1;
        }

        String confFolder = PublicUtil.getProjectConfigFolder();
        String reportPath = confFolder + "/" + "report/PersonProgramExercises" + subProgramCount + "Session.jasper";

//        InputStream inputStream = new FileInputStream(new java.io.File(reportPath));
//        JasperDesign jasperDesign = JRXmlLoader.load(inputStream);
//        JasperReport jasperReport = JasperCompileManager.compileReport(jasperDesign);

        JasperReport jasperReport  = (JasperReport) JRLoader.loadObject(new java.io.File(reportPath));

        Map<String, Object> parameters = new HashMap();
        parameters.put("coachName", "وحید مداحی");
        parameters.put("programDate", DateUtil.getShamsiDate(program.getProgramDate()));
        parameters.put("personName", person.getFirstName() + ' ' + person.getLastName());
        parameters.put("age", program.getPersonAge());
        parameters.put("tall", program.getPersonTall());
        parameters.put("weight", program.getPersonWeight());
        parameters.put("chest", program.getPersonChest());
        parameters.put("waist", program.getPersonWaist());
        parameters.put("abdomen", program.getPersonAbdomen());
        parameters.put("arm", program.getPersonArm());
        parameters.put("forearm", program.getPersonForeArm());
        parameters.put("thigh", program.getPersonThigh());
        parameters.put("shin", program.getPersonShin());
        parameters.put("butt", program.getPersonButt());
        parameters.put("fatPercentage", program.getPersonFatPercentage());
        parameters.put("fatWeight", program.getPersonFatWeight());
        parameters.put("muscleWeight", program.getPersonMuscleWeight());
        parameters.put("score", program.getPersonScore());

        JasperPrint jasperPrint = JasperFillManager.fillReport(jasperReport, parameters,
                new JRBeanCollectionDataSource(exerciseItems));
        return jasperPrint;
    }
}
