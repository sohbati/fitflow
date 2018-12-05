package com.caa.services;

import com.caa.constants.ProgramConstants;
import com.caa.dao.ExerciseDao;
import com.caa.dao.ProgramDao;
import com.caa.dao.ProgramExerciseItemDao;
import com.caa.model.Exercise;
import com.caa.model.Person;
import com.caa.model.Program;
import com.caa.model.ProgramExerciseItem;
import com.caa.modelview.ImageView;
import com.caa.modelview.ProgramExerciseItemView;
import com.caa.modelview.ProgramView;
import com.caa.report.ExportReport;
import com.caa.report.ProgramExercisesReportDTO;
import com.caa.util.DateUtil;
import com.caa.util.ImageUtil;
import org.apache.commons.collections.map.HashedMap;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Repository;
import org.springframework.transaction.annotation.Isolation;
import org.springframework.transaction.annotation.Transactional;
import org.springframework.util.StreamUtils;
import org.springframework.web.bind.annotation.PathVariable;

import javax.imageio.ImageIO;
import java.io.IOException;
import java.util.ArrayList;
import java.util.List;
import java.util.Map;
import java.util.concurrent.ExecutorService;

import static com.caa.constants.ProgramConstants.*;
/**
 * Created by Reza on 27/08/2018.
 */
@Repository
@Transactional(isolation= Isolation.READ_COMMITTED)
public class ProgramService {


    @Autowired
    private ProgramDao programDao;
    @Autowired
    private ExerciseDao exerciseDao;
    @Autowired
    private ProgramExerciseItemDao programExerciseItemDao;

    @Autowired
    private ExerciseService exerciseService;
    @Autowired
    ProgramExerciseItemService programExerciseItemService;

    @Autowired
    PersonService personService;

    public String[][] getPersonAllSizes(long personId){
            String[][] personAllSizes;

        List<Program> programs = programDao.findByPersonId(personId);
        if (programs == null || programs.size() == 0) {
            return new String[0][0];
        }
        personAllSizes = new String[16][programs.size()  +1];
        personAllSizes[0][0] = ROW_NAME_AGE;
        personAllSizes[1][0] = ROW_NAME_TALL;
        personAllSizes[2][0] = ROW_NAME_WEIGHT;
        personAllSizes[3][0] = ROW_NAME_CHEST;
        personAllSizes[4][0] = ROW_NAME_WAIST;
        personAllSizes[5][0] = ROW_NAME_ABDOMEN;
        personAllSizes[6][0] = ROW_NAME_ARM;
        personAllSizes[7][0] = ROW_NAME_FORE_ARM;
        personAllSizes[8][0] = ROW_NAME_THIGH;
        personAllSizes[9][0] = ROW_NAME_SHIN;
        personAllSizes[10][0] = ROW_NAME_BUTT;
        personAllSizes[11][0] = ROW_NAME_FAT_PERCENTAGE;
        personAllSizes[12][0] = ROW_NAME_FAT_WEIGHT;
        personAllSizes[13][0] = ROW_NAME_MUSCLE_WEIGHT;
        personAllSizes[14][0] = ROW_NAME_SCORE;
        personAllSizes[15][0] = ROW_NAME_DATE;

        for (int i = 0; i < programs.size(); i++) {
            personAllSizes[0][i+1] = programs.get(i).getPersonAge() + "";
            personAllSizes[1][i+1] = programs.get(i).getPersonTall() + "";
            personAllSizes[2][i+1] = programs.get(i).getPersonWeight() + "";
            personAllSizes[3][i+1] = programs.get(i).getPersonChest() + "";
            personAllSizes[4][i+1] = programs.get(i).getPersonWaist() + "";
            personAllSizes[5][i+1] = programs.get(i).getPersonAbdomen() + "";
            personAllSizes[6][i+1] = programs.get(i).getPersonArm() + "";
            personAllSizes[7][i+1] = programs.get(i).getPersonForeArm() + "";
            personAllSizes[8][i+1] = programs.get(i).getPersonThigh() + "";
            personAllSizes[9][i+1] = programs.get(i).getPersonShin() + "";
            personAllSizes[10][i+1] = programs.get(i).getPersonButt() + "";
            personAllSizes[11][i+1] = programs.get(i).getPersonFatPercentage() + "";
            personAllSizes[12][i+1] = programs.get(i).getPersonFatWeight() + "";
            personAllSizes[13][i+1] = programs.get(i).getPersonMuscleWeight() + "";
            personAllSizes[14][i+1] = programs.get(i).getPersonScore() + "";
            personAllSizes[15][i+1] = DateUtil.getShamsiDate(programs.get(i).getProgramDate());
        }

        return  personAllSizes;
    }

    public Program findOne(long id) {
        return programDao.findOne(id);
    }

    public ImageView getProgramexerciseImage(long programId) {
        String imageBase64 = "";
        if (programId > 0) {
            Program program = findOne(programId);
            Person person = personService.findOne(program.getPersonId());
            List<ProgramExerciseItemView> viewList = getProgramExerciseList(program);

            ExportReport exportReport = new ExportReport();
            List<ProgramExercisesReportDTO> reportDTOList = exerciseService.convertProgramExerciseToReportDTO(viewList);
            imageBase64 = exportReport.getProgramExerciseAsImage(reportDTOList, program, person);
        }
        ImageView imageView = new ImageView();
        imageView.setContent(imageBase64);
        return imageView;
        //return imageBase64;
    }

    public byte[] getProgramExerciseListAsImage(long programId) throws IOException {
        Program program = findOne(programId);
        Person person = personService.findOne(program.getPersonId());

        List<ProgramExerciseItemView> viewList = getProgramExerciseList(program);
        List<ProgramExercisesReportDTO> reportDTOList = exerciseService.convertProgramExerciseToReportDTO(viewList);
        byte[] bytes = ExportReport.getProgramExerciseAsImageInBytes(reportDTOList, program, person);
        return bytes;
    }

    public byte[] getProgramExerciseListAsPDF(long programId) throws IOException {
        Program program = findOne(programId);
        Person person = personService.findOne(program.getPersonId());

        List<ProgramExerciseItemView> viewList = getProgramExerciseList(program);
        List<ProgramExercisesReportDTO> reportDTOList = exerciseService.convertProgramExerciseToReportDTO(viewList);
        byte[] bytes = ExportReport.getProgramExerciseAsPDFInBytes(reportDTOList, program, person);
        return bytes;
    }

    private List<ProgramExerciseItemView> getProgramExerciseList(Program program) {
        List<ProgramExerciseItem> exerciseItems = programExerciseItemService.findByProgramId(program.getId());
        //TODO improve code
        List<ProgramExerciseItemView> viewList = new ArrayList<>();
        if (exerciseItems != null && exerciseItems.size() > 0) {
            exerciseItems.stream().forEach(programExerciseItem -> {
                ProgramExerciseItemView viewItem = new ProgramExerciseItemView(programExerciseItem);
                Exercise ex = exerciseDao.getOne(programExerciseItem.getExerciseId());
                String code = "(" + ex.getCode() + ")  ";
                viewItem.setExerciseName( code + ex.getName());
                viewList.add(viewItem);
            });
        }
        return viewList;
    }

    public Iterable<ProgramView> getPersonsPrograms(long personId) {
        List<Program> list = programDao.findByPersonId(personId);
        List<ProgramView> result = new ArrayList<>();

        for(Program p : list) {
            ProgramView pv = new ProgramView(p);
            pv.setShamsiProgramDate(DateUtil.getShamsiDate(p.getProgramDate()));
            pv.setPersonName(p.getPerson().getFirstName() + " " + p.getPerson().getLastName()
                    + "(" + p.getPerson().getMobileNumber() + ")");

            List<ProgramExerciseItem> exerciseItems = programExerciseItemDao.findByProgramId(p.getId());
            //TODO improve code
            List<ProgramExerciseItemView> viewList = new ArrayList<>();
            if (exerciseItems != null && exerciseItems.size() > 0) {
                exerciseItems.stream().forEach(programExerciseItem -> {
                    ProgramExerciseItemView viewItem = new ProgramExerciseItemView(programExerciseItem);
                    viewItem.setExerciseName(
                            exerciseDao.getOne(programExerciseItem.getExerciseId()).getName());
                    viewList.add(viewItem);
                });
            }

            viewList.stream().forEach(exerciseItem -> {
                switch (exerciseItem.getSubExerciseId()) {
                    case 1 : pv.getProgramExercise1Items().add(exerciseItem); break;
                    case 2 : pv.getProgramExercise2Items().add(exerciseItem); break;
                    case 3 : pv.getProgramExercise3Items().add(exerciseItem); break;
                    case 4 : pv.getProgramExercise4Items().add(exerciseItem); break;
                    case 5 : pv.getProgramExercise5Items().add(exerciseItem); break;
                    case 6 : pv.getProgramExercise6Items().add(exerciseItem); break;
                }
            });
            result.add(pv);
        }
        return result;
    }

    @Transactional
    public ProgramView saveProgram( ProgramView programView) {

        Program p = getProgramFromView(programView);
        List<ProgramExerciseItemView> items1 = programView.getProgramExercise1Items();
        List<ProgramExerciseItemView> items2 = programView.getProgramExercise2Items();
        List<ProgramExerciseItemView> items3 = programView.getProgramExercise3Items();
        List<ProgramExerciseItemView> items4 = programView.getProgramExercise4Items();
        List<ProgramExerciseItemView> items5 = programView.getProgramExercise5Items();
        List<ProgramExerciseItemView> items6 = programView.getProgramExercise6Items();


        programDao.save(p);

        List<ProgramExerciseItem> allEntityList = new ArrayList<>();
        allEntityList.addAll(getSubExerciseItems(items1 , 1, p.getId()));
        allEntityList.addAll(getSubExerciseItems(items2 , 2, p.getId()));
        allEntityList.addAll(getSubExerciseItems(items3 , 3, p.getId()));
        allEntityList.addAll(getSubExerciseItems(items4 , 4, p.getId()));
        allEntityList.addAll(getSubExerciseItems(items5 , 5, p.getId()));
        allEntityList.addAll(getSubExerciseItems(items6 , 6, p.getId()));

        programExerciseItemDao.deleteByProgramId(p.getId());



        programExerciseItemDao.save(allEntityList);
        programView.setId(p.getId());
        return programView;
    }

    private Program getProgramFromView(ProgramView view) {
        Program program = new Program();
        program.setId(view.getId());
        program.setDescription(view.getDescription());
        program.setPerson(view.getPerson());
        program.setPersonAbdomen(view.getPersonAbdomen());
        program.setPersonAge(view.getPersonAge());
        program.setPersonArm(view.getPersonArm());
        program.setPersonChest(view.getPersonChest());
        program.setPersonFatPercentage(view.getPersonFatPercentage());
        program.setPersonFatWeight(view.getPersonFatWeight());
        program.setPersonForeArm(view.getPersonForeArm());
        program.setPersonId(view.getPerson().getId());
        program.setPersonMuscleWeight(view.getPersonMuscleWeight());
        program.setPersonScore(view.getPersonScore());
        program.setPersonShin(view.getPersonShin());
        program.setPersonButt(view.getPersonButt());
        program.setPersonTall(view.getPersonTall());
        program.setPersonThigh(view.getPersonThigh());
        program.setPersonWaist(view.getPersonWaist());
        program.setPersonWeight(view.getPersonWeight());
        program.setProgramDate(DateUtil.getGregorianDate(view.getShamsiProgramDate()));
        program.setProgramName(view.getProgramName());
        return program;
    }

    private List<ProgramExerciseItem> getSubExerciseItems(List<ProgramExerciseItemView> items,
                                                          int subExerciseIndex, Long pid) {
        if (items != null && items.size() > 0) {
            items.forEach(programExerciseItem -> {
                programExerciseItem.setProgramId(pid);
            });
        }

        List<ProgramExerciseItem> entityList = new ArrayList<>();
        items.stream().forEach(programExerciseItemView ->  {
            ProgramExerciseItem item = new ProgramExerciseItem();
            item.setId(programExerciseItemView.getId());
            item.setProgramId(programExerciseItemView.getProgramId());
            item.setExerciseId(programExerciseItemView.getExerciseId());
            item.setSubExerciseId(subExerciseIndex);
            item.setExerciseSet(programExerciseItemView.getExerciseSet());
            item.setExerciseRepeat(programExerciseItemView.getExerciseRepeat());
            item.setExerciseRepeatType(programExerciseItemView.getExerciseRepeatType());
            item.setDescription(programExerciseItemView.getDescription());
            entityList.add(item);
        });
        return entityList;
    }

    @Transactional
    public void deleteProgram(@PathVariable("id") long id) {
        programExerciseItemDao.deleteByProgramId(id);
        programDao.delete(id);
    }

    public void saveProgramPicture(byte[] img, String personMobileNumber, long programId, String imageName, String type) throws IOException {
        byte[] shrinkedImage = ImageUtil.shrinkImage(img, type);
        ImageUtil.saveToFileSystem(personMobileNumber + "/" + programId, imageName, type, img);
        ImageUtil.saveToFileSystem(personMobileNumber + "/" + programId + "/" + "small", imageName, type, shrinkedImage);
    }

    public byte[] loadProgramPictureOriginal(String personMobileNumber, long programId, int imageNumber) throws IOException {
        String path = personMobileNumber + "/" + programId;
        return loadProgramPicture(path, imageNumber);
    }

    public byte[] loadProgramPictureShrinked(String personMobileNumber, long programId, int imageNumber) throws IOException {
        String path = personMobileNumber + "/" + programId + "/" + "small";
        return loadProgramPicture(path, imageNumber);
    }

     private byte[] loadProgramPicture(String path, int imageNumber) throws IOException {
        byte[] img;
        img = getImage(path,
                    ProgramConstants.PROGRAM_PICTURE_NAME + imageNumber + ".jpg");
        if (img.length == 0) {
            img = getImage(path,
                    ProgramConstants.PROGRAM_PICTURE_NAME + imageNumber + ".png");
        }
        return img;
    }

    private byte[] getImage(String path, String fileName) {
        try {
            byte[] img = ImageUtil.loadImageInbytes(path, fileName);
            if (img != null && img.length > 0) {
                return img;
            }
        }catch (Exception e) {
            // ignore
        }
        return new byte[0];
    }
}