webpackJsonp(["main"],{

/***/ "../../../../../src/$$_lazy_route_resource lazy recursive":
/***/ (function(module, exports) {

function webpackEmptyAsyncContext(req) {
	// Here Promise.resolve().then() is used instead of new Promise() to prevent
	// uncatched exception popping up in devtools
	return Promise.resolve().then(function() {
		throw new Error("Cannot find module '" + req + "'.");
	});
}
webpackEmptyAsyncContext.keys = function() { return []; };
webpackEmptyAsyncContext.resolve = webpackEmptyAsyncContext;
module.exports = webpackEmptyAsyncContext;
webpackEmptyAsyncContext.id = "../../../../../src/$$_lazy_route_resource lazy recursive";

/***/ }),

/***/ "../../../../../src/app/app.component.css":
/***/ (function(module, exports, __webpack_require__) {

exports = module.exports = __webpack_require__("../../../../css-loader/lib/css-base.js")(false);
// imports


// module
exports.push([module.i, "/* :host {\r\n    position: absolute;\r\n    left: 0;\r\n    right: 0;\r\n    bottom: 0;\r\n    top: 0;\r\n    background-color: #eaeef3;\r\n    color: darkslategrey;\r\n    font-family: \"Verdana\", sans-serif;\r\n  } */", ""]);

// exports


/*** EXPORTS FROM exports-loader ***/
module.exports = module.exports.toString();

/***/ }),

/***/ "../../../../../src/app/app.component.html":
/***/ (function(module, exports) {

module.exports = "<router-outlet></router-outlet>\n"

/***/ }),

/***/ "../../../../../src/app/app.component.ts":
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "a", function() { return AppComponent; });
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_0__angular_core__ = __webpack_require__("../../../core/esm5/core.js");
var __decorate = (this && this.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};

var AppComponent = (function () {
    function AppComponent() {
    }
    AppComponent = __decorate([
        Object(__WEBPACK_IMPORTED_MODULE_0__angular_core__["n" /* Component */])({
            selector: 'app-root',
            template: __webpack_require__("../../../../../src/app/app.component.html"),
            styles: [__webpack_require__("../../../../../src/app/app.component.css")]
        })
    ], AppComponent);
    return AppComponent;
}());



/***/ }),

/***/ "../../../../../src/app/app.module.ts":
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "a", function() { return AppModule; });
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_0__angular_platform_browser__ = __webpack_require__("../../../platform-browser/esm5/platform-browser.js");
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_1__angular_core__ = __webpack_require__("../../../core/esm5/core.js");
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_2__angular_router__ = __webpack_require__("../../../router/esm5/router.js");
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_3__angular_common_http__ = __webpack_require__("../../../common/esm5/http.js");
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_4__app_component__ = __webpack_require__("../../../../../src/app/app.component.ts");
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_5__header_header_component__ = __webpack_require__("../../../../../src/app/header/header.component.ts");
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_6__header_nav_nav_component__ = __webpack_require__("../../../../../src/app/header/nav/nav.component.ts");
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_7__angular_forms__ = __webpack_require__("../../../forms/esm5/forms.js");
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_8__header_nav_login_login_component__ = __webpack_require__("../../../../../src/app/header/nav/login/login.component.ts");
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_9__views_home_home_component__ = __webpack_require__("../../../../../src/app/views/home/home.component.ts");
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_10__views_init_init_component__ = __webpack_require__("../../../../../src/app/views/init/init.component.ts");
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_11__main_signup_signup_component__ = __webpack_require__("../../../../../src/app/main/signup/signup.component.ts");
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_12__views_auth_auth_component__ = __webpack_require__("../../../../../src/app/views/auth/auth.component.ts");
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_13__header_nav_menu_menu_component__ = __webpack_require__("../../../../../src/app/header/nav/menu/menu.component.ts");
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_14__views_recipes_recipes_component__ = __webpack_require__("../../../../../src/app/views/recipes/recipes.component.ts");
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_15__main_item_item_component__ = __webpack_require__("../../../../../src/app/main/item/item.component.ts");
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_16__main_recipe_recipe_component__ = __webpack_require__("../../../../../src/app/main/recipe/recipe.component.ts");
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_17__services_main_service__ = __webpack_require__("../../../../../src/app/services/main.service.ts");
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_18__services_auth_service__ = __webpack_require__("../../../../../src/app/services/auth.service.ts");
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_19__header_nav_menu_slidebar_slidebar_component__ = __webpack_require__("../../../../../src/app/header/nav/menu/slidebar/slidebar.component.ts");
var __decorate = (this && this.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};




















var appRoutes = [
    { path: '', component: __WEBPACK_IMPORTED_MODULE_10__views_init_init_component__["a" /* InitComponent */] },
    { path: 'home', component: __WEBPACK_IMPORTED_MODULE_9__views_home_home_component__["a" /* HomeComponent */] },
    { path: 'signup', component: __WEBPACK_IMPORTED_MODULE_12__views_auth_auth_component__["a" /* AuthComponent */] },
    { path: 'recipes', component: __WEBPACK_IMPORTED_MODULE_14__views_recipes_recipes_component__["a" /* RecipesComponent */] }
];
var AppModule = (function () {
    function AppModule() {
    }
    AppModule = __decorate([
        Object(__WEBPACK_IMPORTED_MODULE_1__angular_core__["I" /* NgModule */])({
            declarations: [
                __WEBPACK_IMPORTED_MODULE_4__app_component__["a" /* AppComponent */],
                __WEBPACK_IMPORTED_MODULE_5__header_header_component__["a" /* HeaderComponent */],
                __WEBPACK_IMPORTED_MODULE_6__header_nav_nav_component__["a" /* NavComponent */],
                __WEBPACK_IMPORTED_MODULE_8__header_nav_login_login_component__["a" /* LoginComponent */],
                __WEBPACK_IMPORTED_MODULE_9__views_home_home_component__["a" /* HomeComponent */],
                __WEBPACK_IMPORTED_MODULE_10__views_init_init_component__["a" /* InitComponent */],
                __WEBPACK_IMPORTED_MODULE_11__main_signup_signup_component__["a" /* SignupComponent */],
                __WEBPACK_IMPORTED_MODULE_12__views_auth_auth_component__["a" /* AuthComponent */],
                __WEBPACK_IMPORTED_MODULE_13__header_nav_menu_menu_component__["a" /* MenuComponent */],
                __WEBPACK_IMPORTED_MODULE_14__views_recipes_recipes_component__["a" /* RecipesComponent */],
                __WEBPACK_IMPORTED_MODULE_15__main_item_item_component__["a" /* ItemComponent */],
                __WEBPACK_IMPORTED_MODULE_16__main_recipe_recipe_component__["a" /* RecipeComponent */],
                __WEBPACK_IMPORTED_MODULE_19__header_nav_menu_slidebar_slidebar_component__["a" /* SlidebarComponent */]
            ],
            imports: [
                __WEBPACK_IMPORTED_MODULE_0__angular_platform_browser__["a" /* BrowserModule */],
                __WEBPACK_IMPORTED_MODULE_7__angular_forms__["a" /* FormsModule */],
                __WEBPACK_IMPORTED_MODULE_3__angular_common_http__["b" /* HttpClientModule */],
                __WEBPACK_IMPORTED_MODULE_2__angular_router__["b" /* RouterModule */].forRoot(appRoutes)
            ],
            providers: [
                __WEBPACK_IMPORTED_MODULE_18__services_auth_service__["a" /* AuthService */],
                __WEBPACK_IMPORTED_MODULE_17__services_main_service__["a" /* MainService */]
            ],
            bootstrap: [__WEBPACK_IMPORTED_MODULE_4__app_component__["a" /* AppComponent */]]
        })
    ], AppModule);
    return AppModule;
}());



/***/ }),

/***/ "../../../../../src/app/header/header.component.css":
/***/ (function(module, exports, __webpack_require__) {

exports = module.exports = __webpack_require__("../../../../css-loader/lib/css-base.js")(false);
// imports


// module
exports.push([module.i, ":host {\r\n    height: 10vh;\r\n    background-color: darkslategrey;\r\n    box-shadow: 0 1px 10px darkslategrey;\r\n    color: white;\r\n    display: -webkit-box;\r\n    display: -ms-flexbox;\r\n    display: flex;\r\n    -webkit-box-align: center;\r\n        -ms-flex-align: center;\r\n            align-items: center;\r\n}", ""]);

// exports


/*** EXPORTS FROM exports-loader ***/
module.exports = module.exports.toString();

/***/ }),

/***/ "../../../../../src/app/header/header.component.ts":
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "a", function() { return HeaderComponent; });
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_0__angular_core__ = __webpack_require__("../../../core/esm5/core.js");
var __decorate = (this && this.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};
var __metadata = (this && this.__metadata) || function (k, v) {
    if (typeof Reflect === "object" && typeof Reflect.metadata === "function") return Reflect.metadata(k, v);
};

var HeaderComponent = (function () {
    function HeaderComponent() {
    }
    HeaderComponent.prototype.ngOnInit = function () {
    };
    HeaderComponent = __decorate([
        Object(__WEBPACK_IMPORTED_MODULE_0__angular_core__["n" /* Component */])({
            selector: 'app-header',
            template: "<app-nav></app-nav>",
            styles: [__webpack_require__("../../../../../src/app/header/header.component.css")]
        }),
        __metadata("design:paramtypes", [])
    ], HeaderComponent);
    return HeaderComponent;
}());



/***/ }),

/***/ "../../../../../src/app/header/nav/login/login.component.css":
/***/ (function(module, exports, __webpack_require__) {

exports = module.exports = __webpack_require__("../../../../css-loader/lib/css-base.js")(false);
// imports


// module
exports.push([module.i, "form {\r\n    display: -webkit-box;\r\n    display: -ms-flexbox;\r\n    display: flex;\r\n    -webkit-box-orient: horizontal;\r\n    -webkit-box-direction: normal;\r\n        -ms-flex-flow: row nowrap;\r\n            flex-flow: row nowrap;\r\n}\r\n\r\nform > * {\r\n    margin-left: 1.2rem;\r\n}\r\n\r\n.submit {\t\t\r\n    font-size: 1.1rem;\r\n    color: #fff;\r\n    border-radius: 15px;\r\n    outline: none;\r\n    border: 2px solid #fff;\r\n    width: 7rem;\r\n    height: 2rem;\r\n    background-color: darkslategrey;\r\n    cursor: pointer;\t\t\r\n}\t\t\r\n\t\t\r\n.submit:hover {\t\t\r\n    border-color: #fff;\t\t\r\n    background-color: #fff;\t\t\r\n    color: darkslategrey;\r\n    font-weight: bold;\r\n}\r\n\r\ninput {\r\n    outline-color: #eaeef3;\r\n    padding-left: 1rem;\r\n}\r\n", ""]);

// exports


/*** EXPORTS FROM exports-loader ***/
module.exports = module.exports.toString();

/***/ }),

/***/ "../../../../../src/app/header/nav/login/login.component.html":
/***/ (function(module, exports) {

module.exports = "<form #loginForm=\"ngForm\" (submit)=\"onSubmit()\">\t\t\n  <input type=\"email\" [(ngModel)]=\"creds.name\" name=\"name\" placeholder=\"login\">\t\n  <input type=\"password\"  [(ngModel)]=\"creds.password\" name=\"password\" placeholder=\"password\">\n  <button class=\"submit\" type=\"submit\">Login</button>\t\t\t\n</form>\n  \n"

/***/ }),

/***/ "../../../../../src/app/header/nav/login/login.component.ts":
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "a", function() { return LoginComponent; });
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_0__angular_core__ = __webpack_require__("../../../core/esm5/core.js");
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_1__models_auth__ = __webpack_require__("../../../../../src/app/models/auth.ts");
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_2__services_auth_service__ = __webpack_require__("../../../../../src/app/services/auth.service.ts");
var __decorate = (this && this.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};
var __metadata = (this && this.__metadata) || function (k, v) {
    if (typeof Reflect === "object" && typeof Reflect.metadata === "function") return Reflect.metadata(k, v);
};



var LoginComponent = (function () {
    function LoginComponent(authService) {
        this.authService = authService;
        this.creds = new __WEBPACK_IMPORTED_MODULE_1__models_auth__["a" /* Login */]();
    }
    LoginComponent.prototype.onSubmit = function () {
        this.authService.login(this.creds);
    };
    LoginComponent = __decorate([
        Object(__WEBPACK_IMPORTED_MODULE_0__angular_core__["n" /* Component */])({
            selector: 'app-login',
            template: __webpack_require__("../../../../../src/app/header/nav/login/login.component.html"),
            styles: [__webpack_require__("../../../../../src/app/header/nav/login/login.component.css")]
        }),
        __metadata("design:paramtypes", [__WEBPACK_IMPORTED_MODULE_2__services_auth_service__["a" /* AuthService */]])
    ], LoginComponent);
    return LoginComponent;
}());



/***/ }),

/***/ "../../../../../src/app/header/nav/menu/menu.component.css":
/***/ (function(module, exports, __webpack_require__) {

exports = module.exports = __webpack_require__("../../../../css-loader/lib/css-base.js")(false);
// imports


// module
exports.push([module.i, ".slider {\t\t\r\n    font-size: 1.1rem;\r\n    color: #fff;\r\n    border-radius: 15px;\r\n    outline: none;\r\n    border: 2px solid #fff;\r\n    width: 7rem;\r\n    height: 2rem;\r\n    background-color: darkslategrey;\r\n    cursor: pointer;\t\t\r\n}\t", ""]);

// exports


/*** EXPORTS FROM exports-loader ***/
module.exports = module.exports.toString();

/***/ }),

/***/ "../../../../../src/app/header/nav/menu/menu.component.html":
/***/ (function(module, exports) {

module.exports = "<app-slidebar></app-slidebar>\n<div class=\"slider\" (click)=\"toggleMenu();\">Show</div>"

/***/ }),

/***/ "../../../../../src/app/header/nav/menu/menu.component.ts":
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "a", function() { return MenuComponent; });
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_0__angular_core__ = __webpack_require__("../../../core/esm5/core.js");
var __decorate = (this && this.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};
var __metadata = (this && this.__metadata) || function (k, v) {
    if (typeof Reflect === "object" && typeof Reflect.metadata === "function") return Reflect.metadata(k, v);
};

var MenuComponent = (function () {
    function MenuComponent() {
        this.menuState = 'out';
    }
    MenuComponent.prototype.ngOnInit = function () {
    };
    MenuComponent.prototype.toggleMenu = function () {
        this.menuState = this.menuState === 'out' ? 'in' : 'out';
    };
    MenuComponent = __decorate([
        Object(__WEBPACK_IMPORTED_MODULE_0__angular_core__["n" /* Component */])({
            selector: 'app-menu',
            template: __webpack_require__("../../../../../src/app/header/nav/menu/menu.component.html"),
            styles: [__webpack_require__("../../../../../src/app/header/nav/menu/menu.component.css")]
        }),
        __metadata("design:paramtypes", [])
    ], MenuComponent);
    return MenuComponent;
}());



/***/ }),

/***/ "../../../../../src/app/header/nav/menu/slidebar/slidebar.component.css":
/***/ (function(module, exports, __webpack_require__) {

exports = module.exports = __webpack_require__("../../../../css-loader/lib/css-base.js")(false);
// imports


// module
exports.push([module.i, ":host {\r\n    position: fixed;\r\n    height: 90vh;\r\n    width: 25%;\r\n    left: auto;\r\n    top: 10vh;\r\n    right: 0;\r\n    bottom: 0;\r\n    background-color: rgb(81, 121, 121);\r\n}\r\n\r\n/* .closeBtn {\r\n   color: #ccc;\r\n   text-align: end;\r\n   font-size: 30px;\r\n} */\r\n\r\n.wrapper{\r\n    width: 80%;\r\n    margin: 10% auto;\r\n    display: -webkit-box;\r\n    display: -ms-flexbox;\r\n    display: flex;\r\n    -webkit-box-orient: vertical;\r\n    -webkit-box-direction: normal;\r\n        -ms-flex-direction: column;\r\n            flex-direction: column;\r\n}\r\n\r\nli {\r\n    font-size: 1.5rem;\r\n    padding-bottom: 1.2rem;\r\n}\r\n", ""]);

// exports


/*** EXPORTS FROM exports-loader ***/
module.exports = module.exports.toString();

/***/ }),

/***/ "../../../../../src/app/header/nav/menu/slidebar/slidebar.component.html":
/***/ (function(module, exports) {

module.exports = "<div class=\"wrapper\">\n  <!-- <div class=\"closeBtn\">&times;</div> -->\n  <ul>\n    <li><a routerLink=\"/home\">My products</a></li>\n    <li><a routerLink=\"/recipes\">Recipes</a></li>\n    <li><a href=\"#\" (click)=\"logout();\">Log out</a></li>\n  </ul>\n</div>"

/***/ }),

/***/ "../../../../../src/app/header/nav/menu/slidebar/slidebar.component.ts":
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "a", function() { return SlidebarComponent; });
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_0__angular_core__ = __webpack_require__("../../../core/esm5/core.js");
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_1__services_auth_service__ = __webpack_require__("../../../../../src/app/services/auth.service.ts");
var __decorate = (this && this.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};
var __metadata = (this && this.__metadata) || function (k, v) {
    if (typeof Reflect === "object" && typeof Reflect.metadata === "function") return Reflect.metadata(k, v);
};


var SlidebarComponent = (function () {
    function SlidebarComponent(authService) {
        this.authService = authService;
    }
    SlidebarComponent.prototype.ngOnInit = function () {
    };
    SlidebarComponent.prototype.logout = function () {
    };
    SlidebarComponent = __decorate([
        Object(__WEBPACK_IMPORTED_MODULE_0__angular_core__["n" /* Component */])({
            selector: 'app-slidebar',
            template: __webpack_require__("../../../../../src/app/header/nav/menu/slidebar/slidebar.component.html"),
            styles: [__webpack_require__("../../../../../src/app/header/nav/menu/slidebar/slidebar.component.css")]
        }),
        __metadata("design:paramtypes", [__WEBPACK_IMPORTED_MODULE_1__services_auth_service__["a" /* AuthService */]])
    ], SlidebarComponent);
    return SlidebarComponent;
}());



/***/ }),

/***/ "../../../../../src/app/header/nav/nav.component.css":
/***/ (function(module, exports, __webpack_require__) {

exports = module.exports = __webpack_require__("../../../../css-loader/lib/css-base.js")(false);
// imports


// module
exports.push([module.i, ":host {\r\n    height: 100%;\r\n    width: 80%;\r\n    margin: 0 auto;\r\n    display: -webkit-box;\r\n    display: -ms-flexbox;\r\n    display: flex;\r\n    -webkit-box-pack: justify;\r\n        -ms-flex-pack: justify;\r\n            justify-content: space-between;\r\n    -webkit-box-align: center;\r\n        -ms-flex-align: center;\r\n            align-items: center;\r\n}\r\n\r\n.btn-nav {\r\n    border: 2px solid #fff;\r\n    border-radius: 25px;\r\n    width: 7rem;\r\n    height: 2rem;\r\n    display: -webkit-box;\r\n    display: -ms-flexbox;\r\n    display: flex;\r\n    -webkit-box-pack: center;\r\n        -ms-flex-pack: center;\r\n            justify-content: center;\r\n}\r\n\r\n.btn-nav > span {\r\n    -ms-flex-item-align: center;\r\n        -ms-grid-row-align: center;\r\n        align-self: center;\r\n    font-weight: bold;\r\n    padding: 0 0.5rem;\r\n}\r\n\r\n.btn-nav:hover {\r\n    background-color: #fff;\r\n    color: darkslategrey;\r\n    cursor: pointer;\r\n}", ""]);

// exports


/*** EXPORTS FROM exports-loader ***/
module.exports = module.exports.toString();

/***/ }),

/***/ "../../../../../src/app/header/nav/nav.component.html":
/***/ (function(module, exports) {

module.exports = "<div class=\"logo\">\n  <h1><a href='/'>SmartFridge</a></h1>\n</div>\n\n<app-login *ngIf=\"showLogin\"></app-login>\n<app-menu *ngIf=\"showMenu\"></app-menu>"

/***/ }),

/***/ "../../../../../src/app/header/nav/nav.component.ts":
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "a", function() { return NavComponent; });
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_0__angular_core__ = __webpack_require__("../../../core/esm5/core.js");
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_1__services_auth_service__ = __webpack_require__("../../../../../src/app/services/auth.service.ts");
var __decorate = (this && this.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};
var __metadata = (this && this.__metadata) || function (k, v) {
    if (typeof Reflect === "object" && typeof Reflect.metadata === "function") return Reflect.metadata(k, v);
};


var NavComponent = (function () {
    function NavComponent(authService) {
        this.authService = authService;
        this.showLogin = false;
        this.showMenu = false;
    }
    NavComponent.prototype.ngOnInit = function () {
        var authentificated = this.authService.checkLogin();
        (authentificated) ? (this.showMenu = true) : (this.showLogin = true);
    };
    NavComponent = __decorate([
        Object(__WEBPACK_IMPORTED_MODULE_0__angular_core__["n" /* Component */])({
            selector: 'app-nav',
            template: __webpack_require__("../../../../../src/app/header/nav/nav.component.html"),
            styles: [__webpack_require__("../../../../../src/app/header/nav/nav.component.css")]
        }),
        __metadata("design:paramtypes", [__WEBPACK_IMPORTED_MODULE_1__services_auth_service__["a" /* AuthService */]])
    ], NavComponent);
    return NavComponent;
}());



/***/ }),

/***/ "../../../../../src/app/main/item/item.component.css":
/***/ (function(module, exports, __webpack_require__) {

exports = module.exports = __webpack_require__("../../../../css-loader/lib/css-base.js")(false);
// imports


// module
exports.push([module.i, "", ""]);

// exports


/*** EXPORTS FROM exports-loader ***/
module.exports = module.exports.toString();

/***/ }),

/***/ "../../../../../src/app/main/item/item.component.html":
/***/ (function(module, exports) {

module.exports = "<p>\n  item works!\n</p>\n"

/***/ }),

/***/ "../../../../../src/app/main/item/item.component.ts":
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "a", function() { return ItemComponent; });
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_0__angular_core__ = __webpack_require__("../../../core/esm5/core.js");
var __decorate = (this && this.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};
var __metadata = (this && this.__metadata) || function (k, v) {
    if (typeof Reflect === "object" && typeof Reflect.metadata === "function") return Reflect.metadata(k, v);
};

var ItemComponent = (function () {
    function ItemComponent() {
    }
    ItemComponent.prototype.ngOnInit = function () {
    };
    ItemComponent = __decorate([
        Object(__WEBPACK_IMPORTED_MODULE_0__angular_core__["n" /* Component */])({
            selector: 'app-item',
            template: __webpack_require__("../../../../../src/app/main/item/item.component.html"),
            styles: [__webpack_require__("../../../../../src/app/main/item/item.component.css")]
        }),
        __metadata("design:paramtypes", [])
    ], ItemComponent);
    return ItemComponent;
}());



/***/ }),

/***/ "../../../../../src/app/main/recipe/recipe.component.css":
/***/ (function(module, exports, __webpack_require__) {

exports = module.exports = __webpack_require__("../../../../css-loader/lib/css-base.js")(false);
// imports


// module
exports.push([module.i, "", ""]);

// exports


/*** EXPORTS FROM exports-loader ***/
module.exports = module.exports.toString();

/***/ }),

/***/ "../../../../../src/app/main/recipe/recipe.component.html":
/***/ (function(module, exports) {

module.exports = "<div class=\"wrapper\">\n</div>"

/***/ }),

/***/ "../../../../../src/app/main/recipe/recipe.component.ts":
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "a", function() { return RecipeComponent; });
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_0__angular_core__ = __webpack_require__("../../../core/esm5/core.js");
var __decorate = (this && this.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};
var __metadata = (this && this.__metadata) || function (k, v) {
    if (typeof Reflect === "object" && typeof Reflect.metadata === "function") return Reflect.metadata(k, v);
};

var RecipeComponent = (function () {
    function RecipeComponent() {
    }
    RecipeComponent.prototype.ngOnInit = function () {
    };
    RecipeComponent = __decorate([
        Object(__WEBPACK_IMPORTED_MODULE_0__angular_core__["n" /* Component */])({
            selector: 'app-recipe',
            template: __webpack_require__("../../../../../src/app/main/recipe/recipe.component.html"),
            styles: [__webpack_require__("../../../../../src/app/main/recipe/recipe.component.css")]
        }),
        __metadata("design:paramtypes", [])
    ], RecipeComponent);
    return RecipeComponent;
}());



/***/ }),

/***/ "../../../../../src/app/main/signup/signup.component.css":
/***/ (function(module, exports, __webpack_require__) {

exports = module.exports = __webpack_require__("../../../../css-loader/lib/css-base.js")(false);
// imports


// module
exports.push([module.i, ":host {\r\n    width: 50%;\r\n}\r\n\r\ninput {\t\t\r\n    display: block;\r\n    width: 100%;\r\n    margin-top: 0.5rem;\r\n    height: 1.3rem;\r\n    outline-color: #eaeef3;\r\n    padding-left: 1rem;\r\n}\r\n\t\t\r\nform > * {\t\t\r\n    margin-bottom: 1rem;\t\t\r\n}\t\t\r\n\t\t\r\n.submit {\t\t\r\n    float: right;\r\n    font-size: 1.1rem;\r\n    color: #184646;\r\n    margin-top: 0.5rem;\r\n    border-radius: 15px;\r\n    outline: none;\r\n    border: 2px solid #457d7d;\r\n    width: 7rem;\r\n    height: 2rem;\r\n    background-color: #a7c7c7;\r\n    cursor: pointer;\t\t\r\n}\t\t\r\n\t\t\r\n.submit:hover {\t\t\r\n    border-color: darkslategrey;\t\t\r\n    background-color: darkslategrey;\t\t\r\n    color: white;\t\t\r\n}\t\t\r\n\t\t\r\nheader {\t\t\r\n    height: 8vh;\r\n    margin-bottom: 2rem;\t\t\r\n    display: -webkit-box;\t\t\r\n    display: -ms-flexbox;\t\t\r\n    display: flex;\t\t\r\n    -webkit-box-pack: justify;\t\t\r\n        -ms-flex-pack: justify;\t\t\r\n            justify-content: space-between;\t\t\r\n    background-color: inherit;\t\t\r\n    color: inherit;\t\t\r\n}\t\t\r\n", ""]);

// exports


/*** EXPORTS FROM exports-loader ***/
module.exports = module.exports.toString();

/***/ }),

/***/ "../../../../../src/app/main/signup/signup.component.html":
/***/ (function(module, exports) {

module.exports = "<header>\t\n  <h1>Signup</h1>\t\t\t\n</header>\t\t\n<form #submitForm=\"ngForm\" (submit)=\"onSubmit()\">\t\n  <label for=\"name\">User name</label>\n  <input type=\"text\" id=\"name\" [(ngModel)]=\"user.name\" name=\"name\">\t\n  <label for=\"password\">Password</label>\t\t\n  <input type=\"password\" id=\"password\" [(ngModel)]=\"user.password\" name=\"password\" required>\t\t\n  <label for=\"pwdconfirm\">Confirm password</label>\t\t\n  <input type=\"password\" id=\"pwdconfirm\" [(ngModel)]=\"user.pwdconfirm\" name=\"pwdconfirm\" required>\t\t\n  <label for=\"email\">Email</label>\t\t\n  <input type=\"email\" id=\"email\" [(ngModel)]=\"user.email\" name=\"email\" required>\t\t\n  <button class=\"submit\" type=\"submit\">Signup</button>\t\t\t\n</form>\n"

/***/ }),

/***/ "../../../../../src/app/main/signup/signup.component.ts":
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "a", function() { return SignupComponent; });
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_0__angular_core__ = __webpack_require__("../../../core/esm5/core.js");
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_1__models_auth__ = __webpack_require__("../../../../../src/app/models/auth.ts");
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_2__services_auth_service__ = __webpack_require__("../../../../../src/app/services/auth.service.ts");
var __decorate = (this && this.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};
var __metadata = (this && this.__metadata) || function (k, v) {
    if (typeof Reflect === "object" && typeof Reflect.metadata === "function") return Reflect.metadata(k, v);
};



var SignupComponent = (function () {
    function SignupComponent(authService) {
        this.authService = authService;
        this.user = new __WEBPACK_IMPORTED_MODULE_1__models_auth__["b" /* User */]();
    }
    SignupComponent.prototype.onSubmit = function () {
        this.authService.signup(this.user);
    };
    SignupComponent = __decorate([
        Object(__WEBPACK_IMPORTED_MODULE_0__angular_core__["n" /* Component */])({
            selector: 'app-signup',
            template: __webpack_require__("../../../../../src/app/main/signup/signup.component.html"),
            styles: [__webpack_require__("../../../../../src/app/main/signup/signup.component.css")]
        }),
        __metadata("design:paramtypes", [__WEBPACK_IMPORTED_MODULE_2__services_auth_service__["a" /* AuthService */]])
    ], SignupComponent);
    return SignupComponent;
}());



/***/ }),

/***/ "../../../../../src/app/models/auth.ts":
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "a", function() { return Login; });
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "b", function() { return User; });
var Login = (function () {
    function Login(obj) {
        this.name = obj && obj.name;
    }
    return Login;
}());

var User = (function () {
    function User(obj) {
        this.name = obj && obj.name;
    }
    return User;
}());



/***/ }),

/***/ "../../../../../src/app/services/auth.service.ts":
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "a", function() { return AuthService; });
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_0__angular_core__ = __webpack_require__("../../../core/esm5/core.js");
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_1__angular_router__ = __webpack_require__("../../../router/esm5/router.js");
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_2__angular_common_http__ = __webpack_require__("../../../common/esm5/http.js");
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_3__environments_environment__ = __webpack_require__("../../../../../src/environments/environment.ts");
var __decorate = (this && this.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};
var __metadata = (this && this.__metadata) || function (k, v) {
    if (typeof Reflect === "object" && typeof Reflect.metadata === "function") return Reflect.metadata(k, v);
};




var AuthService = (function () {
    function AuthService(router, http) {
        this.router = router;
        this.http = http;
        this.auth = false;
    }
    AuthService.prototype.getCookies = function () {
        if (document.cookie) {
            this.auth = true;
        }
        console.log(document.cookie);
    };
    AuthService.prototype.checkLogin = function () {
        return this.auth;
    };
    AuthService.prototype.redirect = function () {
        (this.auth) ? this.router.navigate(['/home']) : this.router.navigate(['/signup']);
    };
    AuthService.prototype.login = function (creds) {
        var _this = this;
        var body = JSON.stringify({ login: creds.name, pass: creds.password });
        console.log(body);
        this.http
            .post(__WEBPACK_IMPORTED_MODULE_3__environments_environment__["a" /* environment */].apiURL + 'client/login', body, { observe: 'response', withCredentials: true })
            .subscribe(function (res) {
            console.log(res.headers);
            if (res.status === 200) {
                _this.auth = true;
                _this.router.navigate(['/home']);
            }
        }, function (err) {
            console.log('Something went wrong!', err);
        });
    };
    AuthService.prototype.signup = function (user) {
        var _this = this;
        var body = JSON.stringify({ login: user.name, pass: user.password });
        console.log(body);
        this.http
            .post(__WEBPACK_IMPORTED_MODULE_3__environments_environment__["a" /* environment */].apiURL + 'client/signup', body, { observe: 'response', withCredentials: true })
            .subscribe(function (res) {
            console.log(res);
            if (res.status === 200) {
                _this.auth = true;
                _this.router.navigate(['/home']);
            }
        }, function (err) {
            console.log(err);
        });
    };
    AuthService.prototype.logout = function () {
        var _this = this;
        this.http
            .get(__WEBPACK_IMPORTED_MODULE_3__environments_environment__["a" /* environment */].apiURL + 'client/logout', { observe: 'response' })
            .subscribe(function (res) {
            if (res.status === 200) {
                _this.auth = false;
                _this.router.navigate(['/signup']);
            }
        }, function (err) {
            console.log(err);
        });
        this.router.navigate(['/signup']);
    };
    AuthService = __decorate([
        Object(__WEBPACK_IMPORTED_MODULE_0__angular_core__["A" /* Injectable */])(),
        __metadata("design:paramtypes", [__WEBPACK_IMPORTED_MODULE_1__angular_router__["a" /* Router */],
            __WEBPACK_IMPORTED_MODULE_2__angular_common_http__["a" /* HttpClient */]])
    ], AuthService);
    return AuthService;
}());



/***/ }),

/***/ "../../../../../src/app/services/main.service.ts":
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "a", function() { return MainService; });
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_0__angular_core__ = __webpack_require__("../../../core/esm5/core.js");
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_1__angular_common_http__ = __webpack_require__("../../../common/esm5/http.js");
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_2__environments_environment__ = __webpack_require__("../../../../../src/environments/environment.ts");
var __decorate = (this && this.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};
var __metadata = (this && this.__metadata) || function (k, v) {
    if (typeof Reflect === "object" && typeof Reflect.metadata === "function") return Reflect.metadata(k, v);
};



var MainService = (function () {
    function MainService(http) {
        this.http = http;
    }
    MainService.prototype.getRecipes = function () {
        var _this = this;
        this.http.get(__WEBPACK_IMPORTED_MODULE_2__environments_environment__["a" /* environment */].apiURL + 'client/allRecipes', { withCredentials: true }).subscribe(function (data) {
            console.log(data);
            _this.recipes = data;
        });
    };
    MainService.prototype.getProducts = function () {
        var _this = this;
        this.http.get(__WEBPACK_IMPORTED_MODULE_2__environments_environment__["a" /* environment */].apiURL + 'client/fridgeContent').subscribe(function (data) {
            console.log(data);
            _this.products = data;
        });
    };
    MainService.prototype.showProducts = function () {
        return this.products;
    };
    MainService.prototype.showRecipes = function () {
        return this.recipes;
    };
    MainService = __decorate([
        Object(__WEBPACK_IMPORTED_MODULE_0__angular_core__["A" /* Injectable */])(),
        __metadata("design:paramtypes", [__WEBPACK_IMPORTED_MODULE_1__angular_common_http__["a" /* HttpClient */]])
    ], MainService);
    return MainService;
}());



/***/ }),

/***/ "../../../../../src/app/views/auth/auth.component.css":
/***/ (function(module, exports, __webpack_require__) {

exports = module.exports = __webpack_require__("../../../../css-loader/lib/css-base.js")(false);
// imports


// module
exports.push([module.i, ".wrapper {\r\n    height: 85vh;\r\n    width: 80%;\r\n    margin: 0 auto;\r\n    display: -webkit-box;\r\n    display: -ms-flexbox;\r\n    display: flex;\r\n    -webkit-box-pack: justify;\r\n        -ms-flex-pack: justify;\r\n            justify-content: space-between;\r\n    -webkit-box-align: center;\r\n        -ms-flex-align: center;\r\n            align-items: center;\r\n}\r\n\r\nimg {\r\n    -webkit-box-flex: 1;\r\n        -ms-flex: 1 1 auto;\r\n            flex: 1 1 auto;\r\n    max-width: 40%;\r\n}\r\n\r\n", ""]);

// exports


/*** EXPORTS FROM exports-loader ***/
module.exports = module.exports.toString();

/***/ }),

/***/ "../../../../../src/app/views/auth/auth.component.html":
/***/ (function(module, exports) {

module.exports = "<app-header class=\"header\"></app-header>\n<div class=\"wrapper\">\n  <img src=\"assets/GOPHER_SHARE.png\" alt=\"Gophers wait for you to authorize\">\n  <app-signup></app-signup>\n</div>"

/***/ }),

/***/ "../../../../../src/app/views/auth/auth.component.ts":
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "a", function() { return AuthComponent; });
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_0__angular_core__ = __webpack_require__("../../../core/esm5/core.js");
var __decorate = (this && this.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};
var __metadata = (this && this.__metadata) || function (k, v) {
    if (typeof Reflect === "object" && typeof Reflect.metadata === "function") return Reflect.metadata(k, v);
};

var AuthComponent = (function () {
    function AuthComponent() {
    }
    AuthComponent.prototype.ngOnInit = function () {
    };
    AuthComponent = __decorate([
        Object(__WEBPACK_IMPORTED_MODULE_0__angular_core__["n" /* Component */])({
            selector: 'app-auth',
            template: __webpack_require__("../../../../../src/app/views/auth/auth.component.html"),
            styles: [__webpack_require__("../../../../../src/app/views/auth/auth.component.css")]
        }),
        __metadata("design:paramtypes", [])
    ], AuthComponent);
    return AuthComponent;
}());



/***/ }),

/***/ "../../../../../src/app/views/home/home.component.css":
/***/ (function(module, exports, __webpack_require__) {

exports = module.exports = __webpack_require__("../../../../css-loader/lib/css-base.js")(false);
// imports


// module
exports.push([module.i, "", ""]);

// exports


/*** EXPORTS FROM exports-loader ***/
module.exports = module.exports.toString();

/***/ }),

/***/ "../../../../../src/app/views/home/home.component.html":
/***/ (function(module, exports) {

module.exports = "<app-header></app-header>\r\n"

/***/ }),

/***/ "../../../../../src/app/views/home/home.component.ts":
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "a", function() { return HomeComponent; });
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_0__angular_core__ = __webpack_require__("../../../core/esm5/core.js");
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_1__services_main_service__ = __webpack_require__("../../../../../src/app/services/main.service.ts");
var __decorate = (this && this.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};
var __metadata = (this && this.__metadata) || function (k, v) {
    if (typeof Reflect === "object" && typeof Reflect.metadata === "function") return Reflect.metadata(k, v);
};


var HomeComponent = (function () {
    function HomeComponent(mainService) {
        this.mainService = mainService;
    }
    HomeComponent.prototype.ngOnInit = function () {
        this.mainService.getProducts();
        this.products = this.mainService.showProducts();
    };
    HomeComponent = __decorate([
        Object(__WEBPACK_IMPORTED_MODULE_0__angular_core__["n" /* Component */])({
            selector: 'app-home',
            template: __webpack_require__("../../../../../src/app/views/home/home.component.html"),
            styles: [__webpack_require__("../../../../../src/app/views/home/home.component.css")]
        }),
        __metadata("design:paramtypes", [__WEBPACK_IMPORTED_MODULE_1__services_main_service__["a" /* MainService */]])
    ], HomeComponent);
    return HomeComponent;
}());



/***/ }),

/***/ "../../../../../src/app/views/init/init.component.css":
/***/ (function(module, exports, __webpack_require__) {

exports = module.exports = __webpack_require__("../../../../css-loader/lib/css-base.js")(false);
// imports


// module
exports.push([module.i, ":host {\r\n    height: 100vh;\r\n    background-color: #fff;\r\n    display: -webkit-box;\r\n    display: -ms-flexbox;\r\n    display: flex;\r\n    -webkit-box-orient: vertical;\r\n    -webkit-box-direction: normal;\r\n        -ms-flex-flow: column;\r\n            flex-flow: column;\r\n    -webkit-box-align: center;\r\n        -ms-flex-align: center;\r\n            align-items: center;\r\n    -webkit-box-pack: center;\r\n        -ms-flex-pack: center;\r\n            justify-content: center;\r\n}\r\n\r\nimg {\r\n    width: 20%;\r\n    max-width: 250px;\r\n}", ""]);

// exports


/*** EXPORTS FROM exports-loader ***/
module.exports = module.exports.toString();

/***/ }),

/***/ "../../../../../src/app/views/init/init.component.html":
/***/ (function(module, exports) {

module.exports = "<img src=\"assets/fancy_gopher_renee.jpg\" alt=\"Loading placeholder\">"

/***/ }),

/***/ "../../../../../src/app/views/init/init.component.ts":
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "a", function() { return InitComponent; });
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_0__angular_core__ = __webpack_require__("../../../core/esm5/core.js");
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_1__services_auth_service__ = __webpack_require__("../../../../../src/app/services/auth.service.ts");
var __decorate = (this && this.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};
var __metadata = (this && this.__metadata) || function (k, v) {
    if (typeof Reflect === "object" && typeof Reflect.metadata === "function") return Reflect.metadata(k, v);
};


var InitComponent = (function () {
    function InitComponent(authService) {
        this.authService = authService;
    }
    InitComponent.prototype.ngOnInit = function () {
        this.authService.getCookies();
        this.authService.redirect();
    };
    InitComponent = __decorate([
        Object(__WEBPACK_IMPORTED_MODULE_0__angular_core__["n" /* Component */])({
            selector: 'app-init',
            template: __webpack_require__("../../../../../src/app/views/init/init.component.html"),
            styles: [__webpack_require__("../../../../../src/app/views/init/init.component.css")]
        }),
        __metadata("design:paramtypes", [__WEBPACK_IMPORTED_MODULE_1__services_auth_service__["a" /* AuthService */]])
    ], InitComponent);
    return InitComponent;
}());



/***/ }),

/***/ "../../../../../src/app/views/recipes/recipes.component.css":
/***/ (function(module, exports, __webpack_require__) {

exports = module.exports = __webpack_require__("../../../../css-loader/lib/css-base.js")(false);
// imports


// module
exports.push([module.i, ".wrapper {\r\n    width: 50%;\r\n    margin-left: 10%;\r\n    margin-top: 3%;\r\n    box-shadow: 0 0 5px rgb(23, 43, 43);\r\n    padding: 1rem;\r\n    display: -webkit-box;\r\n    display: -ms-flexbox;\r\n    display: flex;\r\n    -webkit-box-align: center;\r\n        -ms-flex-align: center;\r\n            align-items: center;\r\n}\r\n\r\n.wrapper > * {\r\n    margin: 1rem;\r\n}", ""]);

// exports


/*** EXPORTS FROM exports-loader ***/
module.exports = module.exports.toString();

/***/ }),

/***/ "../../../../../src/app/views/recipes/recipes.component.html":
/***/ (function(module, exports) {

module.exports = "<app-header></app-header>\n<div class=\"wrapper\" *ngFor=\"let recipe of recipes\">\n    <h3>{{ recipe.title }}</h3>\n    <div>{{ recipe.description }}</div>\n    <div class=\"time\">{{ recipe.coockingTimeMin }}</div>\n    <div class=\"complexity\">{{ recipe.complexity }}</div>\n    <div>{{ recipe.ingredients }}</div>\n</div>"

/***/ }),

/***/ "../../../../../src/app/views/recipes/recipes.component.ts":
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "a", function() { return RecipesComponent; });
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_0__angular_core__ = __webpack_require__("../../../core/esm5/core.js");
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_1__services_main_service__ = __webpack_require__("../../../../../src/app/services/main.service.ts");
var __decorate = (this && this.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};
var __metadata = (this && this.__metadata) || function (k, v) {
    if (typeof Reflect === "object" && typeof Reflect.metadata === "function") return Reflect.metadata(k, v);
};


var RecipesComponent = (function () {
    function RecipesComponent(mainService) {
        this.mainService = mainService;
    }
    RecipesComponent.prototype.ngOnInit = function () {
        this.mainService.getRecipes();
        this.recipes = this.mainService.showRecipes();
    };
    RecipesComponent = __decorate([
        Object(__WEBPACK_IMPORTED_MODULE_0__angular_core__["n" /* Component */])({
            selector: 'app-recipes',
            template: __webpack_require__("../../../../../src/app/views/recipes/recipes.component.html"),
            styles: [__webpack_require__("../../../../../src/app/views/recipes/recipes.component.css")]
        }),
        __metadata("design:paramtypes", [__WEBPACK_IMPORTED_MODULE_1__services_main_service__["a" /* MainService */]])
    ], RecipesComponent);
    return RecipesComponent;
}());



/***/ }),

/***/ "../../../../../src/environments/environment.ts":
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "a", function() { return environment; });
// The file contents for the current environment will overwrite these during build.
// The build system defaults to the dev environment which uses `environment.ts`, but if you do
// `ng build --env=prod` then `environment.prod.ts` will be used instead.
// The list of which env maps to which file can be found in `.angular-cli.json`.
var environment = {
    production: false,
    apiURL: 'http://localhost:9000/'
};


/***/ }),

/***/ "../../../../../src/main.ts":
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
Object.defineProperty(__webpack_exports__, "__esModule", { value: true });
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_0__angular_core__ = __webpack_require__("../../../core/esm5/core.js");
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_1__angular_platform_browser_dynamic__ = __webpack_require__("../../../platform-browser-dynamic/esm5/platform-browser-dynamic.js");
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_2__app_app_module__ = __webpack_require__("../../../../../src/app/app.module.ts");
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_3__environments_environment__ = __webpack_require__("../../../../../src/environments/environment.ts");




if (__WEBPACK_IMPORTED_MODULE_3__environments_environment__["a" /* environment */].production) {
    Object(__WEBPACK_IMPORTED_MODULE_0__angular_core__["_13" /* enableProdMode */])();
}
Object(__WEBPACK_IMPORTED_MODULE_1__angular_platform_browser_dynamic__["a" /* platformBrowserDynamic */])().bootstrapModule(__WEBPACK_IMPORTED_MODULE_2__app_app_module__["a" /* AppModule */])
    .catch(function (err) { return console.log(err); });


/***/ }),

/***/ 0:
/***/ (function(module, exports, __webpack_require__) {

module.exports = __webpack_require__("../../../../../src/main.ts");


/***/ })

},[0]);
//# sourceMappingURL=main.bundle.js.map