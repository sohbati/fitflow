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


    public List<ProgramExercisesReportDTO> convertProgramExerciseToReportDTO(List<ProgramExerciseItemView> list) {
        List<ProgramExercisesReportDTO> dtos = new ArrayList<>();
        Map<Long, String> mapList = new HashedMap();
        Map<Long, String> mapDescList = new HashedMap();

        list.forEach(programExerciseItemView -> {
            String s = getExerciseItemString(programExerciseItemView);
            long id = programExerciseItemView.getId();
            String desc = programExerciseItemView.getDescription();
            if (mapList.get(id) == null) {
                mapList.put(id, id + " - " + s);
            }else {
                mapList.replace(id, mapList.get(id) + " " +  s);
            }
            mapDescList.put(id, mapDescList.get(id) == null || "".equals(mapDescList.get(id))?
                    desc : "(" + mapDescList.get(id) + ") " + " (" + desc + ")"  );
        });
        mapList.keySet().stream().forEach(id -> {
            ProgramExercisesReportDTO reportDTO = new ProgramExercisesReportDTO();
            reportDTO.setId(id);
            reportDTO.setExercise(mapList.get(id));
            reportDTO.setDescription(mapDescList.get(id));
            dtos.add(reportDTO);
        });

        dtos.sort((o1, o2) -> {return (o1.getId() > o2.getId() ? 1 : -1);  });
        return  dtos;
    }


    private String getExerciseItemString(ProgramExerciseItemView view) {
        StringBuilder sb = new StringBuilder();
        sb.append(" ( ").
                append(view.getExerciseName()).append("  ").
                append(view.getExerciseSet()).append("*").
                append(view.getExerciseRepeat()).append( "  ")
                .append(getRepTypeStr(view.getExerciseRepeatType())).append(")");
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
