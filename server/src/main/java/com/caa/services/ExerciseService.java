package com.caa.services;

import com.caa.dao.ExerciseDao;
import com.caa.modelview.ProgramExerciseItemView;
import com.caa.report.ProgramExercisesReportDTO;
import com.caa.services.security.impl.CustomUserDetailsService;
import org.apache.commons.collections.map.HashedMap;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Repository;

import java.util.ArrayList;
import java.util.List;
import java.util.Map;
import java.util.TreeMap;

/**
 * Created by Reza on 27/08/2018.
 */
@Repository
public class ExerciseService {

    private static final String REPEAT_TYPE_SECOND = "S";
    private static final String REPEAT_TYPE_MINUTE = "M";
    private static final String REPEAT_TYPE_REPEAT = "T";

    @Autowired
    private ExerciseDao exerciseDao;

    public List<ProgramExercisesReportDTO> convertProgramExerciseToReportDTO(List<ProgramExerciseItemView> list) {
        Map<Integer, Map<Long, String>> exerciseNameMapList = new HashedMap();
        Map<Integer, Map<Long, String>> exerciseRepeatSetMapList = new HashedMap();
//        Map<Integer, Map<Long, String>> mapDescList = new HashedMap();

        list.forEach(programExerciseItemView -> {
            String exerciseName = "";
            exerciseName = programExerciseItemView.getExerciseName();
            String set = programExerciseItemView.getExerciseSet() + "";
            String repeat = programExerciseItemView.getExerciseRepeat() +
                    getExerciseRepeatTypeSymbol(programExerciseItemView.getExerciseRepeatType());
            int subExerciseId = programExerciseItemView.getSubExerciseId();
            long id = programExerciseItemView.getId();
            String desc = programExerciseItemView.getDescription();
            // exercise name string
            if (exerciseNameMapList.get(subExerciseId) == null) {
                exerciseNameMapList.put(subExerciseId, new HashedMap());
            }
            if (exerciseNameMapList.get(subExerciseId).get(id) == null) {
                exerciseNameMapList.get(subExerciseId).put(id, id + " - " + exerciseName);
            }else {
                exerciseNameMapList.get(subExerciseId).replace(id, exerciseNameMapList.get(subExerciseId).get(id) + " + " +  exerciseName);
            }
            //exercise repeat set String
            if (exerciseRepeatSetMapList.get(subExerciseId) == null) {
                exerciseRepeatSetMapList.put(subExerciseId, new HashedMap());
            }
            if (exerciseRepeatSetMapList.get(subExerciseId).get(id) == null) {
                exerciseRepeatSetMapList.get(subExerciseId).put(id, set + ". " + repeat);
            }else {
                exerciseRepeatSetMapList.get(subExerciseId).replace(id, exerciseRepeatSetMapList.get(subExerciseId).get(id) + " + " +  repeat);
            }

        });
        List<ProgramExercisesReportDTO> dtos = new ArrayList<>();

        Map<Integer, TreeMap<Long, String>> sortedList = sortList(exerciseNameMapList);

        for (Integer key : sortedList.keySet()) {
            TreeMap<Long, String> tmap = sortedList.get(key);
            Map<Long, String> exerciseRepeatSetMapItemMap = exerciseRepeatSetMapList.get(key);
            if (tmap.size() > 0) {
                for (Long subKey: tmap.keySet()) {
                    addToPrintDtoList(dtos, key, subKey, tmap.get(subKey), exerciseRepeatSetMapItemMap.get(subKey));
                }
            }
        }
        return  dtos;
    }

    private void addToPrintDtoList( List<ProgramExercisesReportDTO> dtos,
                                    Integer subListIndex, Long id, String exercise, String repeatSet ) {
        boolean found = false;
        for (ProgramExercisesReportDTO item: dtos) {
            if (subListIndex == 1 && (item.getProgram1Id() == null || item.getProgram1Id() == 0)){
                addItem(item, subListIndex, id, exercise, repeatSet);
                found = true;
                break;
            }else
            if (subListIndex == 2 && (item.getProgram2Id() == null || item.getProgram2Id() == 0)){
                addItem(item, subListIndex, id, exercise, repeatSet);
                found = true;
                break;
            }else
            if (subListIndex == 3 && (item.getProgram3Id() == null || item.getProgram3Id() == 0)){
                addItem(item, subListIndex, id, exercise, repeatSet);
                found = true;
                break;
            }else
            if (subListIndex == 4 && (item.getProgram4Id() == null || item.getProgram4Id() == 0)){
                addItem(item, subListIndex, id, exercise, repeatSet);
                found = true;
                break;
            }else
            if (subListIndex == 5 && (item.getProgram5Id() == null || item.getProgram5Id() == 0)){
                addItem(item, subListIndex, id, exercise, repeatSet);
                found = true;
                break;
            }else
            if (subListIndex == 6 && (item.getProgram6Id() == null || item.getProgram6Id() == 0)){
                addItem(item, subListIndex, id, exercise, repeatSet);
                found = true;
                break;
            }
        }

        if (!found) {
            ProgramExercisesReportDTO newItem  = new ProgramExercisesReportDTO();
            dtos.add(newItem);
            newItem.setProgram1Exercise("");
            newItem.setProgram2Exercise("");
            newItem.setProgram3Exercise("");
            newItem.setProgram4Exercise("");
            newItem.setProgram5Exercise("");
            newItem.setProgram6Exercise("");
            addItem(newItem, subListIndex, id, exercise, repeatSet);
        }
    }

    private void addItem(ProgramExercisesReportDTO reportDTOItem,
            Integer subListIndex, Long id, String exercise, String repeatSet) {
        switch (subListIndex) {
            case 1 :
                reportDTOItem.setProgram1Id(id);
                reportDTOItem.setProgram1Exercise(exercise);
                reportDTOItem.setProgram1ExerciseRepeatSet(repeatSet);
                break;
            case 2 :
                reportDTOItem.setProgram2Id(id);
                reportDTOItem.setProgram2Exercise(exercise);
                reportDTOItem.setProgram2ExerciseRepeatSet(repeatSet);
                break;
            case 3 :
                reportDTOItem.setProgram3Id(id);
                reportDTOItem.setProgram3Exercise(exercise);
                reportDTOItem.setProgram3ExerciseRepeatSet(repeatSet);
                break;
            case 4 :
                reportDTOItem.setProgram4Id(id);
                reportDTOItem.setProgram4Exercise(exercise);
                reportDTOItem.setProgram4ExerciseRepeatSet(repeatSet);
                break;
            case 5 :
                reportDTOItem.setProgram5Id(id);
                reportDTOItem.setProgram5Exercise(exercise);
                reportDTOItem.setProgram5ExerciseRepeatSet(repeatSet);
                break;
            case 6 :
                reportDTOItem.setProgram6Id(id);
                reportDTOItem.setProgram6Exercise(exercise);
                reportDTOItem.setProgram6ExerciseRepeatSet(repeatSet);
                break;
        }
    }
    /**
     * Sort by converting to TreeMap
     * @param map
     * @return treeMap structure
     */
    private Map<Integer, TreeMap<Long, String>> sortList(Map<Integer, Map<Long, String>> map) {
        Map<Integer, TreeMap<Long, String>> newMap = new HashedMap();
        for (Integer key : map.keySet()) {
//            for (Map.Entry entry : map.entrySet()) {
            TreeMap<Long, String> treeMap = new TreeMap<Long, String>(map.get(key));
            newMap.put(key, treeMap);
        }
        return newMap;
    }

    private String getExerciseRepeatTypeSymbol(String type) {
        switch (type) {
            case REPEAT_TYPE_SECOND : return "\"";
            case REPEAT_TYPE_MINUTE : return "'";
            case REPEAT_TYPE_REPEAT : return "";
        }
        return "";
    }

    public String getNewCode() {
        String code = "1000";
        String tenant = CustomUserDetailsService.getCurrentUserTenant();
        String c = exerciseDao.queryMaxCodeForTenant(tenant);
        if (c != null) {
            code = c;
        }
        code = (Integer.parseInt(code) + 1) + "";
        return code;
    }
}
