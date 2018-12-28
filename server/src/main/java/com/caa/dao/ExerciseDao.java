package com.caa.dao;

import com.caa.model.Exercise;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;
import org.springframework.data.rest.core.annotation.RepositoryRestResource;
import org.springframework.stereotype.Repository;

import javax.transaction.Transactional;
import java.util.List;

@Repository
@RepositoryRestResource
@Transactional
public interface ExerciseDao extends JpaRepository<Exercise, Long> {

    @Query("SELECT e FROM Exercise e WHERE tenantId=?1 AND (e.name like %?2%)")
    public List<Exercise> queryByNameForTenant(String tenantId, String name);

    @Query("SELECT e FROM Exercise e WHERE tenantId=?1 AND (e.code=?2)")
    public List<Exercise> queryByCodeForTenant(String tenantId, String code);

    @Query("SELECT e FROM Exercise e WHERE tenantId=?1 AND (e.latinName=?2)")
    public List<Exercise> queryByLatinNameForTenant(String tenantId, String latinName);

    @Query("SELECT max(ex.code) FROM Exercise ex where tenantId=?1")
    public String queryMaxCodeForTenant(String tenantId);

    @Query("SELECT e FROM Exercise e WHERE tenantId=?1")
    public List<Exercise> queryAllForTenant(String tenantId);

}