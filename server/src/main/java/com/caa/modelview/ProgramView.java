package com.caa.modelview;

import com.caa.model.Program;
import com.caa.model.ProgramExerciseItem;
import lombok.Data;

import java.util.ArrayList;
import java.util.List;

/**
 * Created by Reza on 05/08/2018.
 */
@Data
public class ProgramView extends Program {

    public ProgramView() {

    }

    public ProgramView(Program p) {
        this.setId(p.getId());
        this.setTenantId(p.getTenantId());
        this.setPerson(p.getPerson());
        this.setProgramDate(p.getProgramDate());
        this.setProgramName(p.getProgramName());
        this.setPersonAge(p.getPersonAge());
        this.setPersonTall(p.getPersonTall());
        this.setPersonWeight(p.getPersonWeight());
        this.setPersonChest(p.getPersonChest());
        this.setPersonWaist(p.getPersonWaist());
        this.setPersonAbdomen(p.getPersonAbdomen());
        this.setPersonArm(p.getPersonArm());
        this.setPersonForeArm(p.getPersonForeArm());
        this.setPersonThigh(p.getPersonThigh());
        this.setPersonShin(p.getPersonShin());
        this.setPersonButt(p.getPersonButt());
        this.setPersonFatPercentage(p.getPersonFatPercentage());
        this.setPersonFatWeight(p.getPersonFatWeight());
        this.setPersonMuscleWeight(p.getPersonMuscleWeight());
        this.setPersonScore(p.getPersonScore());
        this.setDescription(p.getDescription());
    }

    private String shamsiProgramDate;
    private String personName;
    private List<ProgramExerciseItemView> programExercise1Items = new ArrayList<>();
    private List<ProgramExerciseItemView> programExercise2Items = new ArrayList<>();
    private List<ProgramExerciseItemView> programExercise3Items = new ArrayList<>();
    private List<ProgramExerciseItemView> programExercise4Items = new ArrayList<>();
    private List<ProgramExerciseItemView> programExercise5Items = new ArrayList<>();
    private List<ProgramExerciseItemView> programExercise6Items = new ArrayList<>();

}
