import {AbstractControl, Validators} from "angular2/common";
import {ValidatorFn} from "./validator-fn.interface";
import {ValidatorResult} from "./validator-result.interface";
import {isPresent} from "./is-present.function";

export class NumberValidators {
    static between(min: number, max: number): ValidatorFn {
        return (control: AbstractControl): ValidatorResult => {
            if (isPresent(Validators.required(control))) return null;
            let v: number = control.value;
            return isNaN(v) || v < min || v > max ?
                { "between": { "min": min, "max": max } } :
                null;
        };
    };
}