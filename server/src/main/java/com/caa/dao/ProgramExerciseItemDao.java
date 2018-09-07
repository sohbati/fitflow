package com.caa.dao;

import com.caa.model.Exercise;
import com.caa.model.ProgramExerciseItem;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.rest.core.annotation.RepositoryRestResource;
import org.springframework.stereotype.Repository;

import javax.transaction.Transactional;
import java.util.List;

@Repository
@RepositoryRestResource
@Transactional
public interface ProgramExerciseItemDao extends JpaRepository<ProgramExerciseItem, Long> {

    @Transactional
    public Long deleteByProgramId(long programId);

    public List<ProgramExerciseItem> findByProgramId(long id);
    public List<ProgramExerciseItem> findByExerciseId(long id);
}