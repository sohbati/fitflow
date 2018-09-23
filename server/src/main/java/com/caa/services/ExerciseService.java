package com.caa.services;

import com.caa.dao.ExerciseDao;
import com.caa.dao.ProgramDao;
import com.caa.modelview.ProgramExerciseItemView;
import com.caa.report.ProgramExercisesReportDTO;
import org.apache.commons.collections.map.HashedMap;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Repository;

import java.util.ArrayList;
import java.util.List;
import java.util.Map;
import java.util.TreeMap;
import java.util.stream.Collectors;

/**
 * Created by Reza on 27/08/2018.
 */
@Repository
public class ExerciseService {

    private static final String REPEAT_TYPE_SECOND = "S";
    private static final String REPEAT_TYPE_MINUTE = "M";
    private static final String REPEAT_TYPE_REPEAT = "T";

    private static final String REPEAT_TYPE_SECOND_DESC = "ثانیه";
    private static final String REPEAT_TYPE_MINUTE_DESC = "دقیقه";
    private static final String REPEAT_TYPE_REPEAT_DESC = "تکرار";
    @Autowired
    private ExerciseDao exerciseDao;


//    public List<ProgramExercisesReportDTO> convertProgramExerciseToReportDTO(List<ProgramExerciseItemView> list) {
//        List<ProgramExercisesReportDTO> dtos = new ArrayList<>();
//        Map<Long, String> mapList = new HashedMap();
//        Map<Long, String> mapDescList = new HashedMap();
//
//        list.forEach(programExerciseItemView -> {
//            String s = getExerciseItemString(programExerciseItemView);
//            long id = programExerciseItemView.getId();
//            String desc = programExerciseItemView.getDescription();
//            if (mapList.get(id) == null) {
//                mapList.put(id, id + " - " + s);
//            }else {
//                mapList.replace(id, mapList.get(id) + " " +  s);
//            }
//            mapDescList.put(id, mapDescList.get(id) == null || "".equals(mapDescList.get(id))?
//                    desc : "(" + mapDescList.get(id) + ") " + " (" + desc + ")"  );
//        });
//        mapList.keySet().stream().forEach(id -> {
//            ProgramExercisesReportDTO reportDTO = new ProgramExercisesReportDTO();
//            reportDTO.setId(id);
//            reportDTO.setExercise(mapList.get(id));
//            reportDTO.setDescription(mapDescList.get(id));
//            dtos.add(reportDTO);
//        });
//
//        dtos.sort((o1, o2) -> {return (o1.getId() > o2.getId() ? 1 : -1);  });
//        return  dtos;
//    }

    public List<ProgramExercisesReportDTO> convertProgramExerciseToReportDTO(List<ProgramExerciseItemView> list) {
        Map<Integer, Map<Long, String>> mapList = new HashedMap();
//        Map<Integer, Map<Long, String>> mapDescList = new HashedMap();

        list.forEach(programExerciseItemView -> {
            String s = getExerciseItemString(programExerciseItemView);
            int subExerciseId = programExerciseItemView.getSubExerciseId();
            long id = programExerciseItemView.getId();
            String desc = programExerciseItemView.getDescription();
            if (mapList.get(subExerciseId) == null) {
                mapList.put(subExerciseId, new HashedMap());
            }
            if (mapList.get(subExerciseId).get(id) == null) {
                mapList.get(subExerciseId).put(id, id + " - " + s);
            }else {
                mapList.get(subExerciseId).replace(id, mapList.get(subExerciseId).get(id) + " " +  s);
            }
        });
        List<ProgramExercisesReportDTO> dtos = new ArrayList<>();

        Map<Integer, TreeMap<Long, String>> sortedList = sortList(mapList);

        for (Integer key : sortedList.keySet()) {
            TreeMap<Long, String> tmap = sortedList.get(key);
            if (tmap.size() > 0) {
                for (Long subKey: tmap.keySet()) {
                    addToPrintDtoList(dtos, key, subKey, tmap.get(subKey));
                }
            }
        }
        return  dtos;
    }

    private void addToPrintDtoList( List<ProgramExercisesReportDTO> dtos,
                                    Integer subListIndex, Long id, String exercise ) {
        boolean found = false;
        for (ProgramExercisesReportDTO item: dtos) {
            if (subListIndex == 1 && (item.getProgram1Id() == null || item.getProgram1Id() == 0)){
                addItem(item, subListIndex, id, exercise);
                found = true;
            }else
            if (subListIndex == 2 && (item.getProgram2Id() == null || item.getProgram2Id() == 0)){
                addItem(item, subListIndex, id, exercise);
                found = true;
            }else
            if (subListIndex == 3 && (item.getProgram3Id() == null || item.getProgram3Id() == 0)){
                addItem(item, subListIndex, id, exercise);
                found = true;
            }else
            if (subListIndex == 4 && (item.getProgram4Id() == null || item.getProgram4Id() == 0)){
                addItem(item, subListIndex, id, exercise);
                found = true;
            }
        }

        if (!found) {
            ProgramExercisesReportDTO newItem  = new ProgramExercisesReportDTO();
            dtos.add(newItem);
            newItem.setProgram1Exercise("");
            newItem.setProgram2Exercise("");
            newItem.setProgram3Exercise("");
            newItem.setProgram4Exercise("");
            addItem(newItem,subListIndex, id, exercise);
        }
    }

    private void addItem(ProgramExercisesReportDTO reportDTOItem,
            Integer subListIndex, Long id, String exercise) {
        switch (subListIndex) {
            case 1 :
                reportDTOItem.setProgram1Id(id);
                reportDTOItem.setProgram1Exercise(exercise);
                break;
            case 2 :
                reportDTOItem.setProgram2Id(id);
                reportDTOItem.setProgram2Exercise(exercise);
                break;
            case 3 :
                reportDTOItem.setProgram3Id(id);
                reportDTOItem.setProgram3Exercise(exercise);
                break;
            case 4 :
                reportDTOItem.setProgram4Id(id);
                reportDTOItem.setProgram4Exercise(exercise);
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

    private String getExerciseItemString(ProgramExerciseItemView view) {
        StringBuilder sb = new StringBuilder();
        sb.append(" ").
                append(view.getExerciseName()).append("  ").
                append(view.getExerciseSet()).append("*").
                append(view.getExerciseRepeat()).append( "  ")
                .append(getRepTypeStr(view.getExerciseRepeatType())).append(" ");
        return sb.toString();
    }

    private String getRepTypeStr(String type) {
        switch (type) {
            case REPEAT_TYPE_SECOND : return REPEAT_TYPE_SECOND_DESC;
            case REPEAT_TYPE_MINUTE : return REPEAT_TYPE_MINUTE_DESC;
            case REPEAT_TYPE_REPEAT : return REPEAT_TYPE_REPEAT_DESC;
        }
        return "";
    }
}
