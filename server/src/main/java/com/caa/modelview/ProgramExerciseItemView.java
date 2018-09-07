package com.caa.modelview;

import com.caa.model.ProgramExerciseItem;
import lombok.Data;

/**
 * Created by Reza on 18/08/2018.
 */
@Data
public class ProgramExerciseItemView extends ProgramExerciseItem{
    public ProgramExerciseItemView() {

    }

    public ProgramExerciseItemView(ProgramExerciseItem p) {
        //this.setExerciseName(p.getE);
        this.setDescription(p.getDescription());
        this.setExerciseId(p.getExerciseId());
        this.setExerciseRepeat(p.getExerciseRepeat());
        this.setExerciseRepeatType(p.getExerciseRepeatType());
        this.setExerciseSet(p.getExerciseSet());
        this.setId(p.getId());
        this.setProgramId(p.getProgramId());
    }
    private String exerciseName;
}
