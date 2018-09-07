
import {Person} from './Person';
import {Exercise} from './Exercise';

export class ProgramView {
  id: number;
  programName: string;
  person: Person;
  programDate: Date;
  description: string = '';

  programExerciseItems: ProgramExerciseItem[];

  personAge: number;
  personTall: number;
  personWeight: number;
  personChest: number;
  personWaist: number;
  personAbdomen: number;
  personArm: number;
  personForeArm: number;
  personThigh: number;
  personShin: number;
  personFatPercentage: number;
  personFatWeight: number;
  personMuscleWeight: number;
  personScore: number;

  shamsiProgramDate: string;
  personName: string;
}

export class ProgramExerciseItem {
  idpk: number ;
  id: number;
  exerciseId: number;
  exerciseName: string;
  exerciseSet: number;
  exerciseRepeat: number;
  exerciseRepeatType: string;
  exerciseRepeatTypeDesc: string;
}
