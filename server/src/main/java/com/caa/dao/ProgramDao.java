package com.caa.dao;

import com.caa.model.Program;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;
import org.springframework.data.rest.core.annotation.RepositoryRestResource;
import org.springframework.stereotype.Repository;

import javax.transaction.Transactional;
import java.util.List;

@Repository
@RepositoryRestResource
@Transactional
public interface ProgramDao extends JpaRepository<Program, Long> {

    @Query("SELECT e FROM Program e WHERE tenantId=?1 and personId=?2")
    public List<Program> queryByPersonIdForTenant(String tenantId, long personId);

    @Query("SELECT e FROM Program e WHERE tenantId=?1")
    public List<Program> queryAllForTenant(String tenantId);
}