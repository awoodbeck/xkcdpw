import {Component} from "angular2/core";
import {Router, RouteConfig, ROUTER_DIRECTIVES} from "angular2/router";
import {AcrosticComponent} from "./acrostic/acrostic.component";
import {RandomComponent} from "./random/random.component";

@Component({
    selector: "app",
    template: `
    <nav class="navbar navbar-inverse navbar-fixed-top">
        <div class="container">
            <div class="navbar-header">
                <button type="button" class="navbar-toggle collapsed" data-toggle="collapse" data-target="#navbar" aria-expanded="false" aria-controls="navbar">
            <span class="sr-only">Toggle navigation</span>
            <span class="icon-bar"></span>
            <span class="icon-bar"></span>
            <span class="icon-bar"></span>
          </button>
                <a class="navbar-brand" [routerLink]="['Acrostic']">XKCDpw</a>
            </div>
            <div id="navbar" class="collapse navbar-collapse">
                <ul class="nav navbar-nav">
                    <li [class.active]="isActive(['Acrostic'])"><a [routerLink]="['Acrostic']">Acrostic</a></li>
                    <li [class.active]="isActive(['Random'])"><a [routerLink]="['Random']">Random</a></li>
                </ul>
            </div>
        </div>
    </nav>

    <div class="container">
        <router-outlet></router-outlet>
    </div>
    `,
    directives: [ROUTER_DIRECTIVES]
})
@RouteConfig([
    { path: "/acrostic", name: "Acrostic", component: AcrosticComponent, useAsDefault: true },
    { path: "/random", name: "Random", component: RandomComponent },
])
export class AppComponent {
    constructor(private _router: Router) { }

    isActive(name: any[]): boolean {
        return this._router.isRouteActive(this._router.generate(name));
    }
}
