import {AbstractControl} from "angular2/common";

export interface ValidatorFn {
    (c: AbstractControl): { [key: string]: any };
}