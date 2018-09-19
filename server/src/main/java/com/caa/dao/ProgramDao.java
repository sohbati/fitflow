package com.caa.dao;

import com.caa.model.Program;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.rest.core.annotation.RepositoryRestResource;
import org.springframework.stereotype.Repository;

import javax.transaction.Transactional;
import java.util.List;

@Repository
@RepositoryRestResource
@Transactional
public interface ProgramDao extends JpaRepository<Program, Long> {

    public List<Program> findByPersonId(long personId);
}