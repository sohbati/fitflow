package com.caa.services;

import com.caa.dao.ProgramExerciseItemDao;
import com.caa.model.ProgramExerciseItem;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Repository;

import java.util.List;

/**
 * Created by Reza on 19/09/2018.
 */
@Repository
public class ProgramExerciseItemService {
    @Autowired
    ProgramExerciseItemDao programExerciseItemDao;


    public List<ProgramExerciseItem> findByProgramId(long id) {
        return programExerciseItemDao.findByProgramId(id);
    }
}
