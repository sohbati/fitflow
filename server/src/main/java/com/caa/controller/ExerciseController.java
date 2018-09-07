package com.caa.controller;

import com.caa.dao.ExerciseDao;
import com.caa.dao.ProgramExerciseItemDao;
import com.caa.model.Exercise;
import com.caa.model.ProgramExerciseItem;
import com.caa.modelview.ExerciseView;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.dao.DataIntegrityViolationException;
import org.springframework.http.HttpEntity;
import org.springframework.http.HttpHeaders;
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
public class ExerciseController {

    private static Logger logger = LoggerFactory.getLogger(ExerciseController.class);
	
    @Autowired
    private ExerciseDao exerciseDao;
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
    	
    	logger.info("findAll entered...");
    	
		return exerciseDao.findAll();
	}

	@RequestMapping(value="/saveExercise", method = RequestMethod.POST, produces = MediaType.APPLICATION_JSON_VALUE)
	@ResponseBody
	@Transactional
	public Exercise saveExercise(@RequestBody Exercise exercise) {
    	
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

		logger.info("findAll entered...");

		List<Exercise> list = exerciseDao.findAll();
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