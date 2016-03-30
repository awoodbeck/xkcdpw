import {Component} from "angular2/core";
import {Control, ControlGroup, FormBuilder, NgFormModel, NgIf, Validators} from "angular2/common";
import {NumberValidators} from "../validators/number.validators";

@Component({
    selector: "acrostic",
    template: `
    <h5>Generate <a href="https://en.wikipedia.org/wiki/Acrostic">acrostical</a> passphrases from a single word.</h5>
    <form [ngFormModel]="form">
        <div class="form-group">
            <label for="word">Word</label>
            <input ngControl="word" type="text" class="form-control" id="word" placeholder="Enter a word" required>
            <div *ngIf="word.dirty && !word.valid">
                <span class="error-block" *ngIf="word.errors.minlength">Please enter a word with at least {{word.errors.minlength.requiredLength}} characters.</span>
                <span class="error-block" *ngIf="word.errors.maxlength">Please enter a word with at most {{word.errors.maxlength.requiredLength}} characters.</span>
                <span class="error-block" *ngIf="word.errors.pattern">Please enter only alpha characters (A-Z).</span>
            </div>
        </div>
        <div class="form-group">
            <label for="count">Number of passphrases to generate</label>
            <input ngControl="count" type="number" class="form-control" id="count" value="10" required>
            <div *ngIf="count.dirty && !count.valid">
                <span class="error-block" *ngIf="count.errors.between">{{count.errors.between.msg}}</span>
            </div>
        </div>
        <button type="submit" class="btn btn-success">Generate</button>
    </form>
    `,
    directives: [NgFormModel, NgIf]
})
export class AcrosticComponent {
    private word: Control;
    private count: Control;
    private form: ControlGroup;

    constructor(private builder: FormBuilder) {
        this.word = new Control("", Validators.compose([
            Validators.minLength(1),
            Validators.maxLength(10),
            Validators.pattern("^[a-zA-Z]+$")
        ]));
        this.count = new Control(10, NumberValidators.between(1, 100));
        this.form = builder.group({
            word: this.word,
            count: this.count
        });
    }
}