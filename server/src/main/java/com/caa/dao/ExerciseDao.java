package com.caa.dao;

import com.caa.model.Exercise;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.rest.core.annotation.RepositoryRestResource;
import org.springframework.stereotype.Repository;

import javax.transaction.Transactional;

@Repository
@RepositoryRestResource
@Transactional
public interface ExerciseDao extends JpaRepository<Exercise, Long> {
	
}