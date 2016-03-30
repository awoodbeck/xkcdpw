import {Component} from "angular2/core";
import {Control, ControlGroup, FormBuilder, NgFormModel, NgIf} from "angular2/common";
import {NumberValidators} from "../validators/number.validators";

@Component({
    selector: "random",
    template: `
    <h5>Generate random passphrases.</h5>
    <form [ngFormModel]="form">
        <div class="form-group">
            <label for="words">Words per passphrase</label>
            <input ngControl="words" type="number" class="form-control" id="words" required>
            <div *ngIf="words.dirty && !words.valid">
                <span class="error-block" *ngIf="words.errors.between">
                    Please choose a number between {{words.errors.between.min}} and {{words.errors.between.max}}, inclusive.
                </span>
            </div>
        </div>
        <div class="form-group">
            <label for="count">Number of passphrases to generate</label>
            <input ngControl="count" type="number" class="form-control" id="count" required>
            <div *ngIf="count.dirty && !count.valid">
                <span class="error-block" *ngIf="count.errors.between">
                    Please choose a number between {{count.errors.between.min}} and {{count.errors.between.max}}, inclusive.
                </span>
            </div>
        </div>
        <button type="submit" class="btn btn-success">Generate</button>
    </form>
    `,
    directives: [NgFormModel, NgIf]
})
export class RandomComponent {
    private words: Control;
    private count: Control;
    private form: ControlGroup;

    constructor(private builder: FormBuilder) {
        this.words = new Control(4, NumberValidators.between(1, 10));
        this.count = new Control(10, NumberValidators.between(1, 100));
        this.form = builder.group({
            words: this.words,
            count: this.count
        });
    }
}