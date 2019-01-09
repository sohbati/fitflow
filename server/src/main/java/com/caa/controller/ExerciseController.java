package com.caa.controller;

import com.caa.dao.ExerciseDao;
import com.caa.dao.ProgramExerciseItemDao;
import com.caa.model.Exercise;
import com.caa.model.ProgramExerciseItem;
import com.caa.modelview.ExerciseView;
import com.caa.services.ExerciseService;
import com.caa.services.security.impl.CustomUserDetailsService;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.dao.DataIntegrityViolationException;
import org.springframework.http.HttpStatus;
import org.springframework.http.MediaType;
import org.springframework.transaction.annotation.Isolation;
import org.springframework.transaction.annotation.Transactional;
import org.springframework.web.bind.annotation.*;

import java.util.ArrayList;
import java.util.List;

@CrossOrigin(origins = "*")
@RestController
@Transactional(isolation= Isolation.READ_COMMITTED)
@RequestMapping(value = "/api", produces = MediaType.APPLICATION_JSON_VALUE)
public class ExerciseController {

    private static Logger logger = LoggerFactory.getLogger(ExerciseController.class);
	
    @Autowired
    private ExerciseDao exerciseDao;

    @Autowired
    private ExerciseService exerciseService;
    @Autowired
    private ProgramExerciseItemDao programExerciseItemDao;

    @RequestMapping(method = RequestMethod.GET, value = "/getExercise/{id:[\\d]+}", produces = MediaType.APPLICATION_JSON_VALUE)
    @ResponseBody
    public Exercise getExercise(@PathVariable("id") long id) {

    	logger.info("getExercise entered: id= " + id);
    	
    	return exerciseDao.findOne(id);
    }
    
	@RequestMapping(value="/getExercises", method = RequestMethod.GET, produces = MediaType.APPLICATION_JSON_VALUE)
	@ResponseBody
	public Iterable<Exercise> getExercises() {
    	
    	logger.info("queryAllForTenant entered...");
		String tenant = CustomUserDetailsService.getCurrentUserTenant();
		return exerciseDao.queryAllForTenant(tenant);
	}

	@RequestMapping(value="/saveExercise", method = RequestMethod.POST, produces = MediaType.APPLICATION_JSON_VALUE)
	@ResponseBody
	@Transactional
	public Exercise saveExercise(@RequestBody Exercise exercise) {

		exercise.setName(exercise.getName() == null ? "" : exercise.getName().trim());
  	// check repeated name
		String tenant = CustomUserDetailsService.getCurrentUserTenant();
		List<Exercise> list = exerciseDao.queryByNameForTenant(tenant, exercise.getName());
		long currentId = exercise.getId();

		if (list != null && list.size() > 0 && list.get(0).getId() != currentId) {
			throw new RuntimeException("نام حرکت تکراری است");
		}

		// check repeated code
		list = exerciseDao.queryByCodeForTenant(tenant, exercise.getCode());
		if (list != null && list.size() > 0 && list.get(0).getId() != currentId) {
			throw new RuntimeException("کد حرکت تکراری است");
		}
		if (exercise.getCode() != null && exercise.getCode().length() > 0) {
			try {
				Integer.parseInt(exercise.getCode());
			} catch (Exception ex) {
				throw new RuntimeException("لطفا کد حرکت عددی وارد گردد");
			}
		}else {
			exercise.setCode(exerciseService.getNewCode());
		}

		if (exercise.getLatinName() == null) {
			exercise.setLatinName("");
		}
		exercise.setLatinName(exercise.getLatinName().trim());
		// check repeated latin name
		list = exerciseDao.queryByLatinNameForTenant(tenant, exercise.getLatinName());
		if (exercise.getLatinName() != null && exercise.getLatinName().length() > 0 &&
				list != null && list.size() > 0 && list.get(0).getId() != currentId) {
			throw new RuntimeException("نام لاتین حرکت تکراری است");
		}

		exercise.setTenantId(tenant);
    	logger.info("saveExercise entered...");
    	exerciseDao.save(exercise);
		return exercise;
	}

	@RequestMapping(method = RequestMethod.DELETE, value = "/deleteExercise/{id:[\\d]+}", produces = MediaType.APPLICATION_JSON_VALUE)
	@Transactional
	@ResponseBody
	public void deleteExercise(@PathVariable("id") long id) {
		List<ProgramExerciseItem> items = programExerciseItemDao.findByExerciseId(id);
		if(items.size() > 0) {
			throw new DataIntegrityViolationException("این حرکت در برنامه یا برنامه هایی استفاده شده است");
		}
		logger.info("delete entered: id= " + id);
		 exerciseDao.delete(id);
	}

	@RequestMapping(value="/getExerciseShortList", method = RequestMethod.GET, produces = MediaType.APPLICATION_JSON_VALUE)
	@ResponseBody
	public Iterable<ExerciseView> getExerciseShortList() {

		logger.info("queryAllForTenant entered...");

		String tenant = CustomUserDetailsService.getCurrentUserTenant();
		List<Exercise> list = exerciseDao.queryAllForTenant(tenant);
		List<ExerciseView> result = new ArrayList<>();

		for (Exercise p : list) {
			ExerciseView pv = new ExerciseView(p);
			result.add(pv);
		}

		return result;
	}

	// Convert a predefined exception to an HTTP Status code
	@ResponseStatus(value= HttpStatus.EXPECTATION_FAILED,
			reason="Data integrity violation")  //
	@ExceptionHandler(DataIntegrityViolationException.class)
	public void conflict() {
		// Nothing to do
	}
}