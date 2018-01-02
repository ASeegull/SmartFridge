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


var AppComponent = (function () {
    function AppComponent(authService) {
        this.authService = authService;
    }
    AppComponent.prototype.ngOnInit = function () {
        this.authService.checkLogin();
        this.authService.redirect();
    };
    AppComponent = __decorate([
        Object(__WEBPACK_IMPORTED_MODULE_0__angular_core__["Component"])({
            selector: 'app-root',
            template: __webpack_require__("../../../../../src/app/app.component.html"),
            styles: [__webpack_require__("../../../../../src/app/app.component.css")]
        }),
        __metadata("design:paramtypes", [__WEBPACK_IMPORTED_MODULE_1__services_auth_service__["a" /* AuthService */]])
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
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_4__angular_platform_browser_animations__ = __webpack_require__("../../../platform-browser/esm5/animations.js");
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_5_ngx_modal__ = __webpack_require__("../../../../ngx-modal/index.js");
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_5_ngx_modal___default = __webpack_require__.n(__WEBPACK_IMPORTED_MODULE_5_ngx_modal__);
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_6__app_component__ = __webpack_require__("../../../../../src/app/app.component.ts");
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_7__header_header_component__ = __webpack_require__("../../../../../src/app/header/header.component.ts");
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_8__header_nav_nav_component__ = __webpack_require__("../../../../../src/app/header/nav/nav.component.ts");
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_9__angular_forms__ = __webpack_require__("../../../forms/esm5/forms.js");
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_10__header_nav_login_login_component__ = __webpack_require__("../../../../../src/app/header/nav/login/login.component.ts");
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_11__views_home_home_component__ = __webpack_require__("../../../../../src/app/views/home/home.component.ts");
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_12__main_signup_signup_component__ = __webpack_require__("../../../../../src/app/main/signup/signup.component.ts");
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_13__views_auth_auth_component__ = __webpack_require__("../../../../../src/app/views/auth/auth.component.ts");
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_14__header_nav_menu_menu_component__ = __webpack_require__("../../../../../src/app/header/nav/menu/menu.component.ts");
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_15__views_recipes_recipes_component__ = __webpack_require__("../../../../../src/app/views/recipes/recipes.component.ts");
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_16__main_item_item_component__ = __webpack_require__("../../../../../src/app/main/item/item.component.ts");
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_17__main_recipe_recipe_component__ = __webpack_require__("../../../../../src/app/main/recipe/recipe.component.ts");
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_18__services_main_service__ = __webpack_require__("../../../../../src/app/services/main.service.ts");
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_19__services_auth_service__ = __webpack_require__("../../../../../src/app/services/auth.service.ts");
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_20__header_nav_menu_slidebar_slidebar_component__ = __webpack_require__("../../../../../src/app/header/nav/menu/slidebar/slidebar.component.ts");
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_21__views_searchrecipes_searchrecipes_component__ = __webpack_require__("../../../../../src/app/views/searchrecipes/searchrecipes.component.ts");
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_22__services_slidebar_service__ = __webpack_require__("../../../../../src/app/services/slidebar.service.ts");
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_23__views_add_agent_add_agent_component__ = __webpack_require__("../../../../../src/app/views/add-agent/add-agent.component.ts");
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_24__main_modal_modal_component__ = __webpack_require__("../../../../../src/app/main/modal/modal.component.ts");
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_25__main_product_product_component__ = __webpack_require__("../../../../../src/app/main/product/product.component.ts");
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_26__views_help_help_component__ = __webpack_require__("../../../../../src/app/views/help/help.component.ts");
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_27__main_add_product_modal_add_product_modal_component__ = __webpack_require__("../../../../../src/app/main/add-product-modal/add-product-modal.component.ts");
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_28__main_update_product_modal_update_product_modal_component__ = __webpack_require__("../../../../../src/app/main/update-product-modal/update-product-modal.component.ts");
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_29__main_search_modal_search_modal_component__ = __webpack_require__("../../../../../src/app/main/search-modal/search-modal.component.ts");
var __decorate = (this && this.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};






























var appRoutes = [
    { path: 'home', component: __WEBPACK_IMPORTED_MODULE_11__views_home_home_component__["a" /* HomeComponent */] },
    { path: 'signup', component: __WEBPACK_IMPORTED_MODULE_13__views_auth_auth_component__["a" /* AuthComponent */] },
    { path: 'recipes', component: __WEBPACK_IMPORTED_MODULE_15__views_recipes_recipes_component__["a" /* RecipesComponent */] },
    { path: 'searchrecipes', component: __WEBPACK_IMPORTED_MODULE_21__views_searchrecipes_searchrecipes_component__["a" /* SearchrecipesComponent */] },
    { path: 'newAgent', component: __WEBPACK_IMPORTED_MODULE_23__views_add_agent_add_agent_component__["a" /* AddAgentComponent */] },
    { path: 'improveapp', component: __WEBPACK_IMPORTED_MODULE_26__views_help_help_component__["a" /* HelpComponent */] }
];
var AppModule = (function () {
    function AppModule() {
    }
    AppModule = __decorate([
        Object(__WEBPACK_IMPORTED_MODULE_1__angular_core__["NgModule"])({
            declarations: [
                __WEBPACK_IMPORTED_MODULE_6__app_component__["a" /* AppComponent */],
                __WEBPACK_IMPORTED_MODULE_7__header_header_component__["a" /* HeaderComponent */],
                __WEBPACK_IMPORTED_MODULE_8__header_nav_nav_component__["a" /* NavComponent */],
                __WEBPACK_IMPORTED_MODULE_10__header_nav_login_login_component__["a" /* LoginComponent */],
                __WEBPACK_IMPORTED_MODULE_11__views_home_home_component__["a" /* HomeComponent */],
                __WEBPACK_IMPORTED_MODULE_12__main_signup_signup_component__["a" /* SignupComponent */],
                __WEBPACK_IMPORTED_MODULE_13__views_auth_auth_component__["a" /* AuthComponent */],
                __WEBPACK_IMPORTED_MODULE_14__header_nav_menu_menu_component__["a" /* MenuComponent */],
                __WEBPACK_IMPORTED_MODULE_15__views_recipes_recipes_component__["a" /* RecipesComponent */],
                __WEBPACK_IMPORTED_MODULE_16__main_item_item_component__["a" /* ItemComponent */],
                __WEBPACK_IMPORTED_MODULE_17__main_recipe_recipe_component__["a" /* RecipeComponent */],
                __WEBPACK_IMPORTED_MODULE_20__header_nav_menu_slidebar_slidebar_component__["a" /* SlidebarComponent */],
                __WEBPACK_IMPORTED_MODULE_21__views_searchrecipes_searchrecipes_component__["a" /* SearchrecipesComponent */],
                __WEBPACK_IMPORTED_MODULE_23__views_add_agent_add_agent_component__["a" /* AddAgentComponent */],
                __WEBPACK_IMPORTED_MODULE_24__main_modal_modal_component__["a" /* ModalComponent */],
                __WEBPACK_IMPORTED_MODULE_25__main_product_product_component__["a" /* ProductComponent */],
                __WEBPACK_IMPORTED_MODULE_26__views_help_help_component__["a" /* HelpComponent */],
                __WEBPACK_IMPORTED_MODULE_27__main_add_product_modal_add_product_modal_component__["a" /* AddProductModalComponent */],
                __WEBPACK_IMPORTED_MODULE_28__main_update_product_modal_update_product_modal_component__["a" /* UpdateProductModalComponent */],
                __WEBPACK_IMPORTED_MODULE_29__main_search_modal_search_modal_component__["a" /* SearchModalComponent */]
            ],
            imports: [
                __WEBPACK_IMPORTED_MODULE_0__angular_platform_browser__["a" /* BrowserModule */],
                __WEBPACK_IMPORTED_MODULE_5_ngx_modal__["ModalModule"],
                __WEBPACK_IMPORTED_MODULE_9__angular_forms__["a" /* FormsModule */],
                __WEBPACK_IMPORTED_MODULE_3__angular_common_http__["b" /* HttpClientModule */],
                __WEBPACK_IMPORTED_MODULE_4__angular_platform_browser_animations__["a" /* BrowserAnimationsModule */],
                __WEBPACK_IMPORTED_MODULE_2__angular_router__["RouterModule"].forRoot(appRoutes)
            ],
            providers: [
                __WEBPACK_IMPORTED_MODULE_19__services_auth_service__["a" /* AuthService */],
                __WEBPACK_IMPORTED_MODULE_18__services_main_service__["a" /* MainService */],
                __WEBPACK_IMPORTED_MODULE_22__services_slidebar_service__["a" /* SlidebarService */]
            ],
            bootstrap: [__WEBPACK_IMPORTED_MODULE_6__app_component__["a" /* AppComponent */]]
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
exports.push([module.i, ":host {\r\n    position: fixed;\r\n    width: 100%;\r\n    height: 10vh;\r\n    background-color: darkslategrey;\r\n    box-shadow: 0 1px 10px darkslategrey;\r\n    color: white;\r\n    display: -webkit-box;\r\n    display: -ms-flexbox;\r\n    display: flex;\r\n    -webkit-box-align: center;\r\n        -ms-flex-align: center;\r\n            align-items: center;\r\n}", ""]);

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
    HeaderComponent = __decorate([
        Object(__WEBPACK_IMPORTED_MODULE_0__angular_core__["Component"])({
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
        Object(__WEBPACK_IMPORTED_MODULE_0__angular_core__["Component"])({
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
exports.push([module.i, ".slider {\t\t\r\n    width: 3rem;\r\n    height: 1.5rem;\r\n    cursor: pointer;\r\n    display: -webkit-box;\r\n    display: -ms-flexbox;\r\n    display: flex;\r\n    -webkit-box-orient:vertical;\r\n    -webkit-box-direction:normal;\r\n        -ms-flex-direction:column;\r\n            flex-direction:column;\r\n    -webkit-box-pack: justify;\r\n        -ms-flex-pack: justify;\r\n            justify-content: space-between\r\n}\r\n\r\n.el {\r\n    width: 100%;\r\n    height: 3px;\r\n    background-color: #fff;\r\n}", ""]);

// exports


/*** EXPORTS FROM exports-loader ***/
module.exports = module.exports.toString();

/***/ }),

/***/ "../../../../../src/app/header/nav/menu/menu.component.html":
/***/ (function(module, exports) {

module.exports = "<app-slidebar [@slidebarState]=\"slidebarService.state\"></app-slidebar>\n<div class=\"slider\" (click)=\"toggleMenu();\">\n    <div class=\"el\"></div>\n    <div class=\"el\"></div>\n    <div class=\"el\"></div>\n</div>"

/***/ }),

/***/ "../../../../../src/app/header/nav/menu/menu.component.ts":
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "a", function() { return MenuComponent; });
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_0__angular_core__ = __webpack_require__("../../../core/esm5/core.js");
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_1__services_slidebar_service__ = __webpack_require__("../../../../../src/app/services/slidebar.service.ts");
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_2__angular_animations__ = __webpack_require__("../../../animations/esm5/animations.js");
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
    function MenuComponent(slidebarService) {
        this.slidebarService = slidebarService;
    }
    MenuComponent.prototype.toggleMenu = function () {
        this.slidebarService.toggle();
    };
    MenuComponent = __decorate([
        Object(__WEBPACK_IMPORTED_MODULE_0__angular_core__["Component"])({
            selector: 'app-menu',
            template: __webpack_require__("../../../../../src/app/header/nav/menu/menu.component.html"),
            styles: [__webpack_require__("../../../../../src/app/header/nav/menu/menu.component.css")],
            animations: [
                Object(__WEBPACK_IMPORTED_MODULE_2__angular_animations__["j" /* trigger */])('slidebarState', [
                    Object(__WEBPACK_IMPORTED_MODULE_2__angular_animations__["g" /* state */])('open', Object(__WEBPACK_IMPORTED_MODULE_2__angular_animations__["h" /* style */])({
                        transform: 'translateX(0)'
                    })),
                    Object(__WEBPACK_IMPORTED_MODULE_2__angular_animations__["g" /* state */])('hidden', Object(__WEBPACK_IMPORTED_MODULE_2__angular_animations__["h" /* style */])({
                        transform: 'translateX(100%)'
                    })),
                    Object(__WEBPACK_IMPORTED_MODULE_2__angular_animations__["i" /* transition */])('inactive => active', Object(__WEBPACK_IMPORTED_MODULE_2__angular_animations__["e" /* animate */])('400ms ease-in-out')),
                    Object(__WEBPACK_IMPORTED_MODULE_2__angular_animations__["i" /* transition */])('active => inactive', Object(__WEBPACK_IMPORTED_MODULE_2__angular_animations__["e" /* animate */])('400ms ease-in-out'))
                ])
            ]
        }),
        __metadata("design:paramtypes", [__WEBPACK_IMPORTED_MODULE_1__services_slidebar_service__["a" /* SlidebarService */]])
    ], MenuComponent);
    return MenuComponent;
}());



/***/ }),

/***/ "../../../../../src/app/header/nav/menu/slidebar/slidebar.component.css":
/***/ (function(module, exports, __webpack_require__) {

exports = module.exports = __webpack_require__("../../../../css-loader/lib/css-base.js")(false);
// imports


// module
exports.push([module.i, ":host {\r\n    position: fixed;\r\n    height: 100vh;\r\n    width: 25%;\r\n    left: auto;\r\n    right: 0;\r\n    bottom: 0;\r\n    background-color: rgb(81, 121, 121);\r\n}\r\n\r\n.closeBtn {\r\n    width: 2rem;\r\n    margin-bottom: 1.5rem;\r\n    -ms-flex-item-align: end;\r\n        align-self: flex-end;\r\n    color: #ccc;\r\n    font-size: 2.2rem;\r\n    cursor: pointer;\r\n}\r\n\r\n.closeBtn:hover {\r\n    color: #fff;\r\n}\r\n\r\n.wrapper{\r\n    width: 80%;\r\n    margin: 10% auto;\r\n    display: -webkit-box;\r\n    display: -ms-flexbox;\r\n    display: flex;\r\n    -webkit-box-orient: vertical;\r\n    -webkit-box-direction: normal;\r\n        -ms-flex-direction: column;\r\n            flex-direction: column;\r\n}\r\n\r\nli {\r\n    font-size: 1.5rem;\r\n    padding-bottom: 1.2rem;\r\n}\r\n\r\na:hover {\r\n    text-shadow: 0 0 5px #91b0b0;\r\n}", ""]);

// exports


/*** EXPORTS FROM exports-loader ***/
module.exports = module.exports.toString();

/***/ }),

/***/ "../../../../../src/app/header/nav/menu/slidebar/slidebar.component.html":
/***/ (function(module, exports) {

module.exports = "<div class=\"wrapper\">\r\n  <div class=\"closeBtn\" (click)=\"toggleMenu();\">&times;</div>\r\n  <ul>\r\n    <li><a routerLink=\"/home\">My products</a></li>\r\n    <li><a routerLink=\"/recipes\">Recipes</a></li>\r\n    <li><a routerLink=\"/searchrecipes\">Search Recipes</a></li>\r\n    <li><a routerLink=\"/newAgent\">Add Agent</a></li>\r\n    <li><a routerLink=\"/improveapp\">Help Us</a></li>\r\n    <li><a href=\"#\" (click)=\"logout();\">Log out</a></li>\r\n  </ul>\r\n</div>"

/***/ }),

/***/ "../../../../../src/app/header/nav/menu/slidebar/slidebar.component.ts":
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "a", function() { return SlidebarComponent; });
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_0__angular_core__ = __webpack_require__("../../../core/esm5/core.js");
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_1__services_auth_service__ = __webpack_require__("../../../../../src/app/services/auth.service.ts");
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_2__services_slidebar_service__ = __webpack_require__("../../../../../src/app/services/slidebar.service.ts");
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
    function SlidebarComponent(authService, slidebarService) {
        this.authService = authService;
        this.slidebarService = slidebarService;
    }
    SlidebarComponent.prototype.ngOnInit = function () {
    };
    SlidebarComponent.prototype.toggleMenu = function () {
        this.slidebarService.toggle();
    };
    SlidebarComponent.prototype.logout = function () {
        this.authService.logout();
    };
    SlidebarComponent = __decorate([
        Object(__WEBPACK_IMPORTED_MODULE_0__angular_core__["Component"])({
            selector: 'app-slidebar',
            template: __webpack_require__("../../../../../src/app/header/nav/menu/slidebar/slidebar.component.html"),
            styles: [__webpack_require__("../../../../../src/app/header/nav/menu/slidebar/slidebar.component.css")]
        }),
        __metadata("design:paramtypes", [__WEBPACK_IMPORTED_MODULE_1__services_auth_service__["a" /* AuthService */], __WEBPACK_IMPORTED_MODULE_2__services_slidebar_service__["a" /* SlidebarService */]])
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
        Object(__WEBPACK_IMPORTED_MODULE_0__angular_core__["Component"])({
            selector: 'app-nav',
            template: __webpack_require__("../../../../../src/app/header/nav/nav.component.html"),
            styles: [__webpack_require__("../../../../../src/app/header/nav/nav.component.css")]
        }),
        __metadata("design:paramtypes", [__WEBPACK_IMPORTED_MODULE_1__services_auth_service__["a" /* AuthService */]])
    ], NavComponent);
    return NavComponent;
}());



/***/ }),

/***/ "../../../../../src/app/main/add-product-modal/add-product-modal.component.css":
/***/ (function(module, exports, __webpack_require__) {

exports = module.exports = __webpack_require__("../../../../css-loader/lib/css-base.js")(false);
// imports


// module
exports.push([module.i, ".modalCenter {\r\n    position: fixed;\r\n    top: 50%;\r\n    left: 50%;\r\n    -webkit-transform: translate(-50%, -50%);\r\n            transform: translate(-50%, -50%);\r\n    background-color: white;\r\n    box-shadow: 0 0 10px darkslategrey;\r\n    z-index: 10;\r\n    outline: none;\r\n}\r\n\r\nmodal-header h3 {\r\n    padding-left: 1.5rem;\r\n}\r\n\r\nmodal-content form {\r\n    width: 90%;\r\n    margin: 0 auto;\r\n    display: -webkit-box;\r\n    display: -ms-flexbox;\r\n    display: flex;\r\n    -ms-flex-pack: distribute;\r\n        justify-content: space-around;\r\n    -webkit-box-orient: horizontal;\r\n    -webkit-box-direction: normal;\r\n        -ms-flex-flow: row wrap;\r\n            flex-flow: row wrap;\r\n}\r\n\r\nmodal-content label {\r\n    -webkit-box-flex: 1;\r\n        -ms-flex: 1 0 20%;\r\n            flex: 1 0 20%;\r\n    margin: 0.7rem;\r\n}\r\n\r\nmodal-content input {\r\n    -webkit-box-flex: 1;\r\n        -ms-flex: 1 0 60%;\r\n            flex: 1 0 60%;\r\n    margin: 0.7rem;\r\n}\r\n\r\nmodal-footer {\r\n    display: -webkit-box;\r\n    display: -ms-flexbox;\r\n    display: flex;\r\n    -webkit-box-pack: end;\r\n        -ms-flex-pack: end;\r\n            justify-content: flex-end;\r\n}\r\n\r\n.close {\r\n    float: rigth;\r\n    background-color: inherit;\r\n    border: none;\r\n    outline: none;\r\n    font-size: 1.5rem;\r\n    font-weight: bold;\r\n    color: darkslategray;\r\n}\r\n\r\nbutton.close:hover {\r\n    color: rgb(24, 43, 43);\r\n}\r\n\r\n.addProduct {\r\n    font-size: 1.1rem;\r\n    color: #fff;\r\n    border-radius: 50%;\r\n    outline: none;\r\n    border: 2px solid #fff;\r\n    width: 15vh;\r\n    height: 15vh;\r\n    background-color: darkslategrey;\r\n    cursor: pointer;\r\n    text-align: center;\r\n    position: fixed;\r\n    bottom: 10%;\r\n    left: 5%;\r\n    display: -webkit-box;\r\n    display: -ms-flexbox;\r\n    display: flex;\r\n    -webkit-box-pack: center;\r\n        -ms-flex-pack: center;\r\n            justify-content: center;\r\n    -webkit-box-align: center;\r\n        -ms-flex-align: center;\r\n            align-items: center;\r\n    box-shadow: 0 0 5px black;\r\n}\r\n\r\n.addProduct:hover {\r\n    color: #cdd6d6;\r\n    background-color:#653959;\r\n}\r\n", ""]);

// exports


/*** EXPORTS FROM exports-loader ***/
module.exports = module.exports.toString();

/***/ }),

/***/ "../../../../../src/app/main/add-product-modal/add-product-modal.component.html":
/***/ (function(module, exports) {

module.exports = "<div class=\"addProduct\" (click)=\"showModal();\">Add<br />Product</div>\n\n<modal #addProduct (onSubmit)=\"onSubmit()\" class=\"modalCenter\" submitButtonLabel=\"Update Product\">\n  <modal-header>\n      <h3>Update agent</h3>    \n      </modal-header>\n      <modal-content>\n            <form #updateProduct=\"ngForm\">\n                <label for=\"product\">Product</label>\t\n                <input type=\"text\" id=\"product\" [(ngModel)]=\"product.name\" name=\"name\">\n                <label for=\"expires\">Expiration Date</label>\t\n                <input type=\"text\" id=\"expires\" [(ngModel)]=\"product.shelfLife\" name=\"shelfLife\">\n                <label for=\"units\">Measurment Units</label>\t\n                <input type=\"text\" id=\"units\" [(ngModel)]=\"product.units\" name=\"units\">\n                <label for=\"imageURL\">Paste link to image</label>\t\n                <input type=\"text\" id=\"imageURL\" [(ngModel)]=\"product.imageURL\" name=\"imageURL\">\n                <p *ngIf=\"success\"><i class=\"fa fa-check\"></i>Your agent is successfully registered</p>\n                <p *ngIf=\"failed\"><i class=\"fa fa-times\" aria-hidden=\"true\"></i>Failed to send request</p>\n            </form>\n      </modal-content>\n  </modal>"

/***/ }),

/***/ "../../../../../src/app/main/add-product-modal/add-product-modal.component.ts":
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "a", function() { return AddProductModalComponent; });
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_0__angular_core__ = __webpack_require__("../../../core/esm5/core.js");
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_1__services_main_service__ = __webpack_require__("../../../../../src/app/services/main.service.ts");
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_2__models_ingredient__ = __webpack_require__("../../../../../src/app/models/ingredient.ts");
var __decorate = (this && this.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};
var __metadata = (this && this.__metadata) || function (k, v) {
    if (typeof Reflect === "object" && typeof Reflect.metadata === "function") return Reflect.metadata(k, v);
};



var AddProductModalComponent = (function () {
    function AddProductModalComponent(mainService) {
        this.mainService = mainService;
        this.product = new (__WEBPACK_IMPORTED_MODULE_2__models_ingredient__["a" /* Ingredient */]);
        this.success = false;
        this.failed = false;
    }
    AddProductModalComponent.prototype.ngOnInit = function () {
    };
    AddProductModalComponent.prototype.showModal = function () {
        this.addProduct.open();
    };
    AddProductModalComponent.prototype.onSubmit = function () {
        var _this = this;
        console.log(this.product);
        this.mainService.addToProductList(this.product).subscribe(function (res) {
            console.log(res);
            if (res.status === 200) {
                console.log('Your agent is successfully registered');
                _this.success = true;
            }
        }, function (err) {
            console.log(err);
            _this.failed = true;
        });
        console.log(this.success, this.failed);
    };
    __decorate([
        Object(__WEBPACK_IMPORTED_MODULE_0__angular_core__["ViewChild"])('addProduct'),
        __metadata("design:type", Object)
    ], AddProductModalComponent.prototype, "addProduct", void 0);
    AddProductModalComponent = __decorate([
        Object(__WEBPACK_IMPORTED_MODULE_0__angular_core__["Component"])({
            selector: 'app-add-product-modal',
            template: __webpack_require__("../../../../../src/app/main/add-product-modal/add-product-modal.component.html"),
            styles: [__webpack_require__("../../../../../src/app/main/add-product-modal/add-product-modal.component.css")]
        }),
        __metadata("design:paramtypes", [__WEBPACK_IMPORTED_MODULE_1__services_main_service__["a" /* MainService */]])
    ], AddProductModalComponent);
    return AddProductModalComponent;
}());



/***/ }),

/***/ "../../../../../src/app/main/item/item.component.css":
/***/ (function(module, exports, __webpack_require__) {

exports = module.exports = __webpack_require__("../../../../css-loader/lib/css-base.js")(false);
// imports


// module
exports.push([module.i, ".product-image {\r\n    width: 100%;\r\n}\r\n\r\n.button {\r\n    height: 1.3rem;\r\n    cursor: pointer;\r\n    margin-top: 0.3rem;\r\n    padding: 0.3rem;\r\n    text-align: center;\r\n    border: 2px solid darkslategrey;\r\n}\r\n\r\n.update {\r\n    background-color: #aab9b9;\r\n}\r\n\r\n.update:hover {\r\n    background-color: darkslategray;\r\n    color: white;\r\n}\r\n\r\n.delete {\r\n    color: #653959;\r\n    background-color: #e4c3c3;\r\n}\r\n\r\n.delete:hover {\r\n    background-color: #653959;\r\n    color: #e4c3c3;\r\n}\r\n", ""]);

// exports


/*** EXPORTS FROM exports-loader ***/
module.exports = module.exports.toString();

/***/ }),

/***/ "../../../../../src/app/main/item/item.component.html":
/***/ (function(module, exports) {

module.exports = "<header>\n  <h3 class=\"title\">{{ product.name | titlecase }}</h3>\n</header>\n<main>\n  <img class=\"product-image\" [src]=\"product.imageURL\" alt=\"Product image\">\n  <p> Shelf Life: {{ product.shelfLife }} </p>\n  <p> Measurement Units: {{ product.units }} </p>\n  <app-update-product-modal [product]=\"product\"></app-update-product-modal>\n  <div class=\"delete button\" (click)=\"deleteProduct()\">Delete Product</div>\n</main>"

/***/ }),

/***/ "../../../../../src/app/main/item/item.component.ts":
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "a", function() { return ItemComponent; });
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_0__angular_core__ = __webpack_require__("../../../core/esm5/core.js");
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_1__models_ingredient__ = __webpack_require__("../../../../../src/app/models/ingredient.ts");
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_2__services_main_service__ = __webpack_require__("../../../../../src/app/services/main.service.ts");
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
    function ItemComponent(mainService) {
        this.mainService = mainService;
    }
    ItemComponent.prototype.updateProduct = function () {
        this.mainService.updateProductList(this.product).subscribe(function (data) { return console.log(data); });
    };
    ItemComponent.prototype.deleteProduct = function () {
        this.mainService.deleteProduct(this.product.ID).subscribe(function (data) { console.log(data); }, function (err) { console.log(err); });
    };
    __decorate([
        Object(__WEBPACK_IMPORTED_MODULE_0__angular_core__["Input"])(),
        __metadata("design:type", __WEBPACK_IMPORTED_MODULE_1__models_ingredient__["a" /* Ingredient */])
    ], ItemComponent.prototype, "product", void 0);
    ItemComponent = __decorate([
        Object(__WEBPACK_IMPORTED_MODULE_0__angular_core__["Component"])({
            selector: 'app-item',
            template: __webpack_require__("../../../../../src/app/main/item/item.component.html"),
            styles: [__webpack_require__("../../../../../src/app/main/item/item.component.css")]
        }),
        __metadata("design:paramtypes", [__WEBPACK_IMPORTED_MODULE_2__services_main_service__["a" /* MainService */]])
    ], ItemComponent);
    return ItemComponent;
}());



/***/ }),

/***/ "../../../../../src/app/main/modal/modal.component.css":
/***/ (function(module, exports, __webpack_require__) {

exports = module.exports = __webpack_require__("../../../../css-loader/lib/css-base.js")(false);
// imports


// module
exports.push([module.i, ".edit {\r\n    margin-right: 0.5rem;\r\n    cursor: pointer;\r\n}\r\n\r\n.edit:hover {\r\n    color: #7ea1a1;\r\n    -webkit-transform: translateY(1px);\r\n            transform: translateY(1px);\r\n}\r\n\r\n.modalCenter {\r\n    position: fixed;\r\n    top: 50%;\r\n    left: 50%;\r\n    -webkit-transform: translate(-50%, -50%);\r\n            transform: translate(-50%, -50%);\r\n    background-color: white;\r\n    box-shadow: 0 0 10px darkslategrey;\r\n    z-index: 10;\r\n    outline: none;\r\n}\r\n\r\nmodal-header h3 {\r\n    padding-left: 1.5rem;\r\n}\r\n\r\nmodal-content form {\r\n    width: 90%;\r\n    margin: 0 auto;\r\n    display: -webkit-box;\r\n    display: -ms-flexbox;\r\n    display: flex;\r\n    -ms-flex-pack: distribute;\r\n        justify-content: space-around;\r\n    -webkit-box-orient: horizontal;\r\n    -webkit-box-direction: normal;\r\n        -ms-flex-flow: row wrap;\r\n            flex-flow: row wrap;\r\n}\r\n\r\nmodal-content label {\r\n    -webkit-box-flex: 1;\r\n        -ms-flex: 1 0 20%;\r\n            flex: 1 0 20%;\r\n    margin: 0.7rem;\r\n}\r\n\r\nmodal-content input {\r\n    -webkit-box-flex: 1;\r\n        -ms-flex: 1 0 60%;\r\n            flex: 1 0 60%;\r\n    margin: 0.7rem;\r\n}\r\n\r\nmodal-footer {\r\n    display: -webkit-box;\r\n    display: -ms-flexbox;\r\n    display: flex;\r\n    -webkit-box-pack: end;\r\n        -ms-flex-pack: end;\r\n            justify-content: flex-end;\r\n}\r\n\r\n.close {\r\n    float: rigth;\r\n    background-color: inherit;\r\n    border: none;\r\n    outline: none;\r\n    font-size: 1.5rem;\r\n    font-weight: bold;\r\n    color: darkslategray;\r\n}\r\n\r\nbutton.close:hover {\r\n    color: rgb(24, 43, 43);\r\n}", ""]);

// exports


/*** EXPORTS FROM exports-loader ***/
module.exports = module.exports.toString();

/***/ }),

/***/ "../../../../../src/app/main/modal/modal.component.html":
/***/ (function(module, exports) {

module.exports = "<div class=\"edit\" title=\"Update agent content\" (click)=\"openModal();\">\n    <i class=\"fa fa-cog\"></i>\n</div>\n\n<modal #editAgent (onSubmit)=\"onSubmit()\" class=\"modalCenter\" submitButtonLabel=\"Update Agent\">\n<modal-header>\n    <h3>Update agent</h3>    \n    </modal-header>\n    <modal-content>\n        <form #updateAgent=\"ngForm\">\n            <label for=\"product\">Product</label>\t\n            <input type=\"text\" id=\"product\" [(ngModel)]=\"agent.product\" name=\"product\">\n            <label for=\"expires\">Expiration Date</label>\t\n            <input type=\"text\" id=\"expires\" [(ngModel)]=\"agent.stateExpires\" name=\"stateExpires\">\n            <p *ngIf=\"success\"><i class=\"fa fa-check\"></i>Your agent is successfully registered</p>\n            <p *ngIf=\"failed\"><i class=\"fa fa-times\" aria-hidden=\"true\"></i>Failed to send request</p>\n        </form>\n    </modal-content>\n</modal>"

/***/ }),

/***/ "../../../../../src/app/main/modal/modal.component.ts":
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "a", function() { return ModalComponent; });
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_0__angular_core__ = __webpack_require__("../../../core/esm5/core.js");
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_1__services_main_service__ = __webpack_require__("../../../../../src/app/services/main.service.ts");
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_2__models_agent__ = __webpack_require__("../../../../../src/app/models/agent.ts");
var __decorate = (this && this.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};
var __metadata = (this && this.__metadata) || function (k, v) {
    if (typeof Reflect === "object" && typeof Reflect.metadata === "function") return Reflect.metadata(k, v);
};



var ModalComponent = (function () {
    function ModalComponent(mainService) {
        this.mainService = mainService;
        this.success = false;
        this.failed = false;
    }
    ModalComponent.prototype.openModal = function () {
        console.log(this.editAgent);
        this.editAgent.open();
    };
    ModalComponent.prototype.onSubmit = function () {
        var _this = this;
        console.log('Hi!');
        this.mainService.updateAgent(this.agent)
            .subscribe(function (res) {
            console.log(res);
            if (res.status === 200) {
                console.log('Your agent is successfully registered');
                _this.success = true;
            }
        }, function (err) {
            console.log(err);
            _this.failed = true;
        });
    };
    __decorate([
        Object(__WEBPACK_IMPORTED_MODULE_0__angular_core__["Input"])(),
        __metadata("design:type", __WEBPACK_IMPORTED_MODULE_2__models_agent__["a" /* Agent */])
    ], ModalComponent.prototype, "agent", void 0);
    __decorate([
        Object(__WEBPACK_IMPORTED_MODULE_0__angular_core__["ViewChild"])('editAgent'),
        __metadata("design:type", Object)
    ], ModalComponent.prototype, "editAgent", void 0);
    ModalComponent = __decorate([
        Object(__WEBPACK_IMPORTED_MODULE_0__angular_core__["Component"])({
            selector: 'app-modal',
            template: __webpack_require__("../../../../../src/app/main/modal/modal.component.html"),
            styles: [__webpack_require__("../../../../../src/app/main/modal/modal.component.css")]
        }),
        __metadata("design:paramtypes", [__WEBPACK_IMPORTED_MODULE_1__services_main_service__["a" /* MainService */]])
    ], ModalComponent);
    return ModalComponent;
}());



/***/ }),

/***/ "../../../../../src/app/main/product/product.component.css":
/***/ (function(module, exports, __webpack_require__) {

exports = module.exports = __webpack_require__("../../../../css-loader/lib/css-base.js")(false);
// imports


// module
exports.push([module.i, ":host > * {\r\n    margin: 1rem;\r\n}\r\n\r\nheader {\r\n    display: -webkit-box;\r\n    display: -ms-flexbox;\r\n    display: flex;\r\n    -webkit-box-pack: justify;\r\n        -ms-flex-pack: justify;\r\n            justify-content: space-between;\r\n}\r\n\r\n.title {\r\n    -webkit-box-flex: 1;\r\n        -ms-flex: 1 1 60%;\r\n            flex: 1 1 60%;\r\n}\r\n\r\n.options {\r\n    display: -webkit-box;\r\n    display: -ms-flexbox;\r\n    display: flex;\r\n    -webkit-box-align: center;\r\n        -ms-flex-align: center;\r\n            align-items: center;\r\n    -webkit-box-flex: 1;\r\n        -ms-flex: 1 0 auto;\r\n            flex: 1 0 auto;\r\n    -webkit-box-pack: end;\r\n        -ms-flex-pack: end;\r\n            justify-content: flex-end;\r\n    font-size: 1.2rem;\r\n}\r\n\r\n.search {\r\n    margin-right: 0.5rem;\r\n    cursor: pointer;\r\n}\r\n\r\n.search:hover {\r\n    color: #7ea1a1;\r\n    -webkit-transform: translateY(1px);\r\n            transform: translateY(1px);\r\n}\r\n\r\n.search-accented {\r\n    text-shadow: 0 0 7px #008686;\r\n}\r\n\r\n.product-image {\r\n    width: 100%;\r\n}\r\n\r\n.product-expired {\r\n    border: 2px solid rgb(216, 45, 45);\r\n}\r\n\r\n.text-ok {\r\n    color: #84d511;\r\n}\r\n\r\n.text-warn {\r\n    color:  rgb(233, 233, 33);\r\n}\r\n\r\n.text-alert {\r\n    color: rgb(216, 45, 45);\r\n}\r\n\r\n.noAgents {\r\n    display: -webkit-box;\r\n    display: -ms-flexbox;\r\n    display: flex;\r\n    -webkit-box-align: center;\r\n        -ms-flex-align: center;\r\n            align-items: center;\r\n    -webkit-box-pack: space-evenly;\r\n        -ms-flex-pack: space-evenly;\r\n            justify-content: space-evenly;\r\n    height: 52vh;\r\n    -webkit-box-orient: vertical;\r\n    -webkit-box-direction: normal;\r\n        -ms-flex-direction: column;\r\n            flex-direction: column;\r\n}\r\n\r\n.addAgent {\r\n    font-size: 1.1rem;\r\n    color: #fff;\r\n    border-radius: 30px;\r\n    outline: none;\r\n    border: 2px solid #fff;\r\n    width: 20%;\r\n    height: 8%;\r\n    padding: 1rem;\r\n    background-color: darkslategrey;\r\n    cursor: pointer;\r\n    text-align: center;\r\n}\r\n\r\n.addAgent:hover {\r\n    color: #cdd6d6;\r\n    background-color: #367373;\r\n}\r\n", ""]);

// exports


/*** EXPORTS FROM exports-loader ***/
module.exports = module.exports.toString();

/***/ }),

/***/ "../../../../../src/app/main/product/product.component.html":
/***/ (function(module, exports) {

module.exports = "<header>\n  <h3 class=\"title\">{{ product.product | titlecase }}</h3>\n  <div class=\"options\">\n      <app-search-modal title=\"Search recepies by this product\"\n          [ngClass]=\"{'search-accented':product.condition === 'warn'}\"\n          [productName]=\"product.product\">\n      </app-search-modal>\n      <app-modal [agent]=\"product\"></app-modal>\n  </div>\n</header>\n<main>\n  <img class=\"product-image\" [src]=\"product.imageURL\" alt=\"product-image\"\n  [ngClass]=\"{'product-expired':product.condition === 'expired'}\">\n  <p>Weight: {{ product.weight }}g</p>\n  <p>\n      <i *ngIf=\"product.condition === 'ok'\" class=\"fa fa-smile-o text-ok\" aria-hidden=\"true\"></i>\n      <i *ngIf=\"product.condition === 'warn'\" class=\"fa fa-exclamation-triangle text-warn\" aria-hidden=\"true\"></i>\n      <i *ngIf=\"product.condition === 'expired'\" class=\"fa fa-frown-o text-alert\" aria-hidden=\"true\"></i>\n      Expires: {{ product.stateExpires }}\n  </p>\n  <p *ngIf=\"product.condition === 'warn'\" class=\"text-warn\">This product is about to expire... Let's see, how we can use it</p>\n  <p *ngIf=\"product.condition === 'expired'\" class=\"text-alert\">Oops... This product isn't safe anymore</p>\n</main>"

/***/ }),

/***/ "../../../../../src/app/main/product/product.component.ts":
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "a", function() { return ProductComponent; });
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_0__angular_core__ = __webpack_require__("../../../core/esm5/core.js");
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_1__models_agent__ = __webpack_require__("../../../../../src/app/models/agent.ts");
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_2__services_main_service__ = __webpack_require__("../../../../../src/app/services/main.service.ts");
var __decorate = (this && this.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};
var __metadata = (this && this.__metadata) || function (k, v) {
    if (typeof Reflect === "object" && typeof Reflect.metadata === "function") return Reflect.metadata(k, v);
};



var ProductComponent = (function () {
    function ProductComponent(mainService) {
        this.mainService = mainService;
    }
    ProductComponent.prototype.ngOnInit = function () {
        console.log(this.product);
    };
    __decorate([
        Object(__WEBPACK_IMPORTED_MODULE_0__angular_core__["Input"])(),
        __metadata("design:type", __WEBPACK_IMPORTED_MODULE_1__models_agent__["a" /* Agent */])
    ], ProductComponent.prototype, "product", void 0);
    ProductComponent = __decorate([
        Object(__WEBPACK_IMPORTED_MODULE_0__angular_core__["Component"])({
            selector: 'app-product',
            template: __webpack_require__("../../../../../src/app/main/product/product.component.html"),
            styles: [__webpack_require__("../../../../../src/app/main/product/product.component.css")]
        }),
        __metadata("design:paramtypes", [__WEBPACK_IMPORTED_MODULE_2__services_main_service__["a" /* MainService */]])
    ], ProductComponent);
    return ProductComponent;
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
        Object(__WEBPACK_IMPORTED_MODULE_0__angular_core__["Component"])({
            selector: 'app-recipe',
            template: __webpack_require__("../../../../../src/app/main/recipe/recipe.component.html"),
            styles: [__webpack_require__("../../../../../src/app/main/recipe/recipe.component.css")]
        }),
        __metadata("design:paramtypes", [])
    ], RecipeComponent);
    return RecipeComponent;
}());



/***/ }),

/***/ "../../../../../src/app/main/search-modal/search-modal.component.css":
/***/ (function(module, exports, __webpack_require__) {

exports = module.exports = __webpack_require__("../../../../css-loader/lib/css-base.js")(false);
// imports


// module
exports.push([module.i, ".card {\r\n    width: 450px;\r\n    box-shadow: 0 0 5px rgb(23, 43, 43);\r\n    padding: 2rem 1rem;\r\n    background-color: white;\r\n    margin: 2rem 1rem;\r\n}\r\n\r\n.card > * {\r\n    margin: 1rem;\r\n}\r\n\r\n.content {\r\n    padding: 13vh 10vw;\r\n    display: -webkit-box;\r\n    display: -ms-flexbox;\r\n    display: flex;\r\n    -webkit-box-orient: horizontal;\r\n    -webkit-box-direction: normal;\r\n        -ms-flex-flow: row wrap;\r\n            flex-flow: row wrap;\r\n    -webkit-box-pack: justify;\r\n        -ms-flex-pack: justify;\r\n            justify-content: space-between;\r\n}\r\n\r\nheader {\r\n    display: -webkit-box;\r\n    display: -ms-flexbox;\r\n    display: flex;\r\n    -webkit-box-pack: justify;\r\n        -ms-flex-pack: justify;\r\n            justify-content: space-between;\r\n}\r\n\r\n.title {\r\n    -webkit-box-flex: 1;\r\n        -ms-flex: 1 1 60%;\r\n            flex: 1 1 60%;\r\n}\r\n\r\n.tags {\r\n    display: -webkit-box;\r\n    display: -ms-flexbox;\r\n    display: flex;\r\n    -webkit-box-align: center;\r\n        -ms-flex-align: center;\r\n            align-items: center;\r\n    -webkit-box-flex: 1;\r\n        -ms-flex: 1 0 auto;\r\n            flex: 1 0 auto;\r\n    -webkit-box-pack: justify;\r\n        -ms-flex-pack: justify;\r\n            justify-content: space-between;\r\n}\r\n\r\n.complexity {\r\n    padding: 0.3rem 0.3rem;\r\n    border-radius: 15px;\r\n    -ms-flex-line-pack: center;\r\n        align-content: center;\r\n}\r\n\r\n.complexity-easy {\r\n    background-color: #84d511;\r\n}\r\n\r\n.complexity-normal {\r\n    background-color: rgb(255, 255, 43);\r\n    color: #4d4141;\r\n}\r\n\r\n.complexity-hard {\r\n    background-color: rgb(216, 45, 45);\r\n    color: #ffe2e2;\r\n}\r\n\r\n.description {\r\n    margin-bottom: 0.7rem;\r\n}", ""]);

// exports


/*** EXPORTS FROM exports-loader ***/
module.exports = module.exports.toString();

/***/ }),

/***/ "../../../../../src/app/main/search-modal/search-modal.component.html":
/***/ (function(module, exports) {

module.exports = "<div  class=\"search\" (click)=\"search();\">\n  <i class=\"fa fa-search\"></i>\n</div>\n\n<modal #searchresults>\n  <modal-header>\n    <h3>Search Results</h3>\n  </modal-header>\n  <modal-content>\n      <article class=\"card\" *ngFor=\"let recipe of recipes\">\n          <header>\n              <h3 class=\"title\">{{ recipe.title }}</h3>\n              <div class=\"tags\">\n                  <div class=\"time\">\n                      <i class=\"fa fa-clock-o\" aria-hidden=\"true\"></i>\n                      {{ recipe.coockingTimeMin }} min</div>\n                  <div class=\"complexity\" [ngClass]=\"{\n                      'complexity-easy':recipe.complexity === 'easy',\n                      'complexity-normal':recipe.complexity === 'normal',\n                      'complexity-hard':recipe.complexity === 'hard'\n                  }\">\n                      <i class=\"fa\" [ngClass]=\"{\n                          'fa-check':recipe.complexity === 'easy',\n                          'fa-cog':recipe.complexity === 'normal',\n                          'fa-cogs':recipe.complexity === 'hard'\n                      }\" arria-hidden=\"true\"></i>\n                  {{ recipe.complexity }}\n                  </div>\n              </div>\n          </header>\n          <main>\n              <div class=\"description\">{{ recipe.description }}</div>\n              <div class=\"ingredients\">{{ recipe.ingredients }}</div>\n          </main>\n      </article>\n  </modal-content>\n</modal>"

/***/ }),

/***/ "../../../../../src/app/main/search-modal/search-modal.component.ts":
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "a", function() { return SearchModalComponent; });
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


var SearchModalComponent = (function () {
    function SearchModalComponent(mainService) {
        this.mainService = mainService;
    }
    SearchModalComponent.prototype.ngOnInit = function () {
    };
    SearchModalComponent.prototype.search = function () {
        var _this = this;
        this.mainService.getRecipesByProduct(this.productName).subscribe(function (data) { _this.recipes = data; }, function (err) { console.log(err); });
        this.searchresults.open();
    };
    __decorate([
        Object(__WEBPACK_IMPORTED_MODULE_0__angular_core__["Input"])(),
        __metadata("design:type", String)
    ], SearchModalComponent.prototype, "productName", void 0);
    __decorate([
        Object(__WEBPACK_IMPORTED_MODULE_0__angular_core__["ViewChild"])('searchresults'),
        __metadata("design:type", Object)
    ], SearchModalComponent.prototype, "searchresults", void 0);
    SearchModalComponent = __decorate([
        Object(__WEBPACK_IMPORTED_MODULE_0__angular_core__["Component"])({
            selector: 'app-search-modal',
            template: __webpack_require__("../../../../../src/app/main/search-modal/search-modal.component.html"),
            styles: [__webpack_require__("../../../../../src/app/main/search-modal/search-modal.component.css")]
        }),
        __metadata("design:paramtypes", [__WEBPACK_IMPORTED_MODULE_1__services_main_service__["a" /* MainService */]])
    ], SearchModalComponent);
    return SearchModalComponent;
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
        Object(__WEBPACK_IMPORTED_MODULE_0__angular_core__["Component"])({
            selector: 'app-signup',
            template: __webpack_require__("../../../../../src/app/main/signup/signup.component.html"),
            styles: [__webpack_require__("../../../../../src/app/main/signup/signup.component.css")]
        }),
        __metadata("design:paramtypes", [__WEBPACK_IMPORTED_MODULE_2__services_auth_service__["a" /* AuthService */]])
    ], SignupComponent);
    return SignupComponent;
}());



/***/ }),

/***/ "../../../../../src/app/main/update-product-modal/update-product-modal.component.css":
/***/ (function(module, exports, __webpack_require__) {

exports = module.exports = __webpack_require__("../../../../css-loader/lib/css-base.js")(false);
// imports


// module
exports.push([module.i, ".button {\r\n    height: 1.3rem;\r\n    cursor: pointer;\r\n    margin-top: 0.3rem;\r\n    padding: 0.3rem;\r\n    text-align: center;\r\n    border: 2px solid darkslategrey;\r\n}\r\n\r\n.update {\r\n    background-color: #aab9b9;\r\n}\r\n\r\n.update:hover {\r\n    background-color: darkslategray;\r\n    color: white;\r\n}\r\n\r\n.modalCenter {\r\n    position: fixed;\r\n    top: 50%;\r\n    left: 50%;\r\n    -webkit-transform: translate(-50%, -50%);\r\n            transform: translate(-50%, -50%);\r\n    background-color: white;\r\n    box-shadow: 0 0 10px darkslategrey;\r\n    z-index: 10;\r\n    outline: none;\r\n}\r\n\r\nmodal-header h3 {\r\n    padding-left: 1.5rem;\r\n}\r\n\r\nmodal-content form {\r\n    width: 90%;\r\n    margin: 0 auto;\r\n    display: -webkit-box;\r\n    display: -ms-flexbox;\r\n    display: flex;\r\n    -ms-flex-pack: distribute;\r\n        justify-content: space-around;\r\n    -webkit-box-orient: horizontal;\r\n    -webkit-box-direction: normal;\r\n        -ms-flex-flow: row wrap;\r\n            flex-flow: row wrap;\r\n}\r\n\r\nmodal-content label {\r\n    -webkit-box-flex: 1;\r\n        -ms-flex: 1 0 20%;\r\n            flex: 1 0 20%;\r\n    margin: 0.7rem;\r\n}\r\n\r\nmodal-content input {\r\n    -webkit-box-flex: 1;\r\n        -ms-flex: 1 0 60%;\r\n            flex: 1 0 60%;\r\n    margin: 0.7rem;\r\n}\r\n\r\nmodal-footer {\r\n    display: -webkit-box;\r\n    display: -ms-flexbox;\r\n    display: flex;\r\n    -webkit-box-pack: end;\r\n        -ms-flex-pack: end;\r\n            justify-content: flex-end;\r\n}\r\n\r\n.close {\r\n    float: rigth;\r\n    background-color: inherit;\r\n    border: none;\r\n    outline: none;\r\n    font-size: 1.5rem;\r\n    font-weight: bold;\r\n    color: darkslategray;\r\n}\r\n", ""]);

// exports


/*** EXPORTS FROM exports-loader ***/
module.exports = module.exports.toString();

/***/ }),

/***/ "../../../../../src/app/main/update-product-modal/update-product-modal.component.html":
/***/ (function(module, exports) {

module.exports = "<div class=\"update button\" (click)=\"openModal()\">Update Product</div>\n\n<modal #updateProduct (onSubmit)=\"onSubmit()\" class=\"modalCenter\" submitButtonLabel=\"Update Product\">\n  <modal-header>\n      <h3>Update agent</h3>    \n      </modal-header>\n      <modal-content>\n            <form #updateProduct=\"ngForm\">\n                <label for=\"product\">Product</label>\t\n                <input type=\"text\" id=\"product\" [(ngModel)]=\"product.name\" name=\"name\">\n                <label for=\"expires\">Expiration Date</label>\t\n                <input type=\"text\" id=\"expires\" [(ngModel)]=\"product.shelfLife\" name=\"shelfLife\">\n                <label for=\"units\">Measurment Units</label>\t\n                <input type=\"text\" id=\"units\" [(ngModel)]=\"product.units\" name=\"units\">\n                <label for=\"imageURL\">Paste link to image</label>\t\n                <input type=\"text\" id=\"imageURL\" [(ngModel)]=\"product.imageURL\" name=\"imageURL\">\n                <p *ngIf=\"success\"><i class=\"fa fa-check\"></i>Your agent is successfully registered</p>\n                <p *ngIf=\"failed\"><i class=\"fa fa-times\" aria-hidden=\"true\"></i>Failed to send request</p>\n            </form>\n      </modal-content>\n  </modal>"

/***/ }),

/***/ "../../../../../src/app/main/update-product-modal/update-product-modal.component.ts":
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "a", function() { return UpdateProductModalComponent; });
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_0__angular_core__ = __webpack_require__("../../../core/esm5/core.js");
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_1__services_main_service__ = __webpack_require__("../../../../../src/app/services/main.service.ts");
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_2__models_ingredient__ = __webpack_require__("../../../../../src/app/models/ingredient.ts");
var __decorate = (this && this.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};
var __metadata = (this && this.__metadata) || function (k, v) {
    if (typeof Reflect === "object" && typeof Reflect.metadata === "function") return Reflect.metadata(k, v);
};



var UpdateProductModalComponent = (function () {
    function UpdateProductModalComponent(mainService) {
        this.mainService = mainService;
        this.success = false;
        this.failed = false;
    }
    UpdateProductModalComponent.prototype.ngOnInit = function () {
    };
    UpdateProductModalComponent.prototype.openModal = function () {
        this.updateProduct.open();
    };
    UpdateProductModalComponent.prototype.onSubmit = function () {
        var _this = this;
        console.log(this.product);
        this.mainService.updateProductList(this.product).subscribe(function (res) {
            console.log(res);
            if (res.status === 200) {
                console.log('Your agent is successfully registered');
                _this.success = true;
            }
        }, function (err) {
            console.log(err);
            _this.failed = true;
        });
    };
    __decorate([
        Object(__WEBPACK_IMPORTED_MODULE_0__angular_core__["Input"])(),
        __metadata("design:type", __WEBPACK_IMPORTED_MODULE_2__models_ingredient__["a" /* Ingredient */])
    ], UpdateProductModalComponent.prototype, "product", void 0);
    __decorate([
        Object(__WEBPACK_IMPORTED_MODULE_0__angular_core__["ViewChild"])('updateProduct'),
        __metadata("design:type", Object)
    ], UpdateProductModalComponent.prototype, "updateProduct", void 0);
    UpdateProductModalComponent = __decorate([
        Object(__WEBPACK_IMPORTED_MODULE_0__angular_core__["Component"])({
            selector: 'app-update-product-modal',
            template: __webpack_require__("../../../../../src/app/main/update-product-modal/update-product-modal.component.html"),
            styles: [__webpack_require__("../../../../../src/app/main/update-product-modal/update-product-modal.component.css")]
        }),
        __metadata("design:paramtypes", [__WEBPACK_IMPORTED_MODULE_1__services_main_service__["a" /* MainService */]])
    ], UpdateProductModalComponent);
    return UpdateProductModalComponent;
}());



/***/ }),

/***/ "../../../../../src/app/models/agent.ts":
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "a", function() { return Agent; });
var Agent = (function () {
    function Agent() {
        // this.agentID = obj && obj.agentID;
        // this.product = obj && obj.product;
        // this.weight = obj && obj.weight;
        // this.stateExpires = obj && obj.stateExpires;
        // this.condition = obj && obj.condition;
    }
    return Agent;
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

/***/ "../../../../../src/app/models/ingredient.ts":
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "a", function() { return Ingredient; });
var Ingredient = (function () {
    function Ingredient() {
    }
    return Ingredient;
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
    AuthService.prototype.checkLogin = function () {
        if (document.cookie) {
            this.auth = true;
        }
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
            .get(__WEBPACK_IMPORTED_MODULE_3__environments_environment__["a" /* environment */].apiURL + 'client/logout', { observe: 'response', withCredentials: true })
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
        Object(__WEBPACK_IMPORTED_MODULE_0__angular_core__["Injectable"])(),
        __metadata("design:paramtypes", [__WEBPACK_IMPORTED_MODULE_1__angular_router__["Router"],
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
        return this.http.get(__WEBPACK_IMPORTED_MODULE_2__environments_environment__["a" /* environment */].apiURL + 'client/allRecipes', { withCredentials: true });
    };
    MainService.prototype.getMyRecipes = function () {
        return this.http.get(__WEBPACK_IMPORTED_MODULE_2__environments_environment__["a" /* environment */].apiURL + 'client/searchRecipes', { withCredentials: true });
    };
    MainService.prototype.getRecipesByProduct = function (name) {
        return this.http.get(__WEBPACK_IMPORTED_MODULE_2__environments_environment__["a" /* environment */].apiURL + ("client/recipes/getByProductName/" + name), { withCredentials: true });
    };
    MainService.prototype.getProducts = function () {
        return this.http.get(__WEBPACK_IMPORTED_MODULE_2__environments_environment__["a" /* environment */].apiURL + 'client/fridgeContent', { withCredentials: true });
    };
    MainService.prototype.addAgent = function (agent) {
        var body = JSON.stringify(agent);
        console.log(body);
        return this.http
            .post(__WEBPACK_IMPORTED_MODULE_2__environments_environment__["a" /* environment */].apiURL + 'client/addAgent', body, { observe: 'response', withCredentials: true });
    };
    MainService.prototype.updateAgent = function (agent) {
        var body = JSON.stringify(agent);
        console.log(body);
        return this.http
            .post(__WEBPACK_IMPORTED_MODULE_2__environments_environment__["a" /* environment */].apiURL + 'client/updateAgent', body, { observe: 'response', withCredentials: true });
    };
    MainService.prototype.getProductList = function () {
        return this.http.get(__WEBPACK_IMPORTED_MODULE_2__environments_environment__["a" /* environment */].apiURL + 'client/getProducts', { withCredentials: true });
    };
    MainService.prototype.addToProductList = function (ingredient) {
        var body = JSON.stringify(ingredient);
        console.log(body);
        return this.http
            .post(__WEBPACK_IMPORTED_MODULE_2__environments_environment__["a" /* environment */].apiURL + 'client/addProduct', body, { observe: 'response', withCredentials: true });
    };
    MainService.prototype.updateProductList = function (ingredient) {
        var body = JSON.stringify(ingredient);
        console.log(body);
        return this.http
            .put(__WEBPACK_IMPORTED_MODULE_2__environments_environment__["a" /* environment */].apiURL + 'client/updateProduct', body, { observe: 'response', withCredentials: true });
    };
    MainService.prototype.deleteProduct = function (id) {
        return this.http
            .delete(__WEBPACK_IMPORTED_MODULE_2__environments_environment__["a" /* environment */].apiURL + ("products/remove/" + id), { observe: 'response', withCredentials: true });
    };
    MainService = __decorate([
        Object(__WEBPACK_IMPORTED_MODULE_0__angular_core__["Injectable"])(),
        __metadata("design:paramtypes", [__WEBPACK_IMPORTED_MODULE_1__angular_common_http__["a" /* HttpClient */]])
    ], MainService);
    return MainService;
}());



/***/ }),

/***/ "../../../../../src/app/services/slidebar.service.ts":
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "a", function() { return SlidebarService; });
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

var SlidebarService = (function () {
    function SlidebarService() {
        this.state = 'hidden';
    }
    SlidebarService.prototype.toggle = function () {
        this.state = this.state === 'hidden' ? 'open' : 'hidden';
    };
    SlidebarService = __decorate([
        Object(__WEBPACK_IMPORTED_MODULE_0__angular_core__["Injectable"])(),
        __metadata("design:paramtypes", [])
    ], SlidebarService);
    return SlidebarService;
}());



/***/ }),

/***/ "../../../../../src/app/views/add-agent/add-agent.component.css":
/***/ (function(module, exports, __webpack_require__) {

exports = module.exports = __webpack_require__("../../../../css-loader/lib/css-base.js")(false);
// imports


// module
exports.push([module.i, ".wrapper {\r\n    width: 100%;\r\n    height: 100vh;\r\n    display: -webkit-box;\r\n    display: -ms-flexbox;\r\n    display: flex;\r\n    -webkit-box-pack: center;\r\n        -ms-flex-pack: center;\r\n            justify-content: center;\r\n    -webkit-box-align: center;\r\n        -ms-flex-align: center;\r\n            align-items: center;\r\n}\r\n\r\n.content {\r\n    width: 70%;\r\n}\r\n\r\ninput {\t\t\r\n    display: block;\r\n    width: 100%;\r\n    margin-top: 0.5rem;\r\n    height: 1.3rem;\r\n    outline-color: #eaeef3;\r\n    padding-left: 1rem;\r\n}\r\n\t\t\r\nform > * {\t\t\r\n    margin-bottom: 1rem;\t\t\r\n}\r\n\r\nfooter {\r\n    display: -webkit-box;\r\n    display: -ms-flexbox;\r\n    display: flex;\r\n    -webkit-box-pack: justify;\r\n        -ms-flex-pack: justify;\r\n            justify-content: space-between;\r\n}\r\n\r\n.submit {\r\n    font-size: 1.1rem;\r\n    color: #184646;\r\n    margin-top: 0.5rem;\r\n    border-radius: 15px;\r\n    outline: none;\r\n    border: 2px solid #457d7d;\r\n    width: 7rem;\r\n    height: 2rem;\r\n    background-color: #a7c7c7;\r\n    cursor: pointer;\t\t\r\n}\r\n\r\n.submit:hover {\t\t\r\n    border-color: darkslategrey;\t\t\r\n    background-color: darkslategrey;\t\t\r\n    color: white;\t\t\r\n}", ""]);

// exports


/*** EXPORTS FROM exports-loader ***/
module.exports = module.exports.toString();

/***/ }),

/***/ "../../../../../src/app/views/add-agent/add-agent.component.html":
/***/ (function(module, exports) {

module.exports = "<app-header></app-header>\n<div class=\"wrapper\">\n  <div class=\"content\">\n<header>\t\n  <h1>New Agent</h1>\t\t\t\n</header>\t\t\n<form #newAgent=\"ngForm\" (submit)=\"onSubmit()\">\t\n  <label for=\"agentID\">Please enter serial number</label>\n  <input type=\"text\" id=\"agentID\" [(ngModel)]=\"agent.agentID\" name=\"agentID\" required>\n  <p>If container is empty yet, you can skip these fields</p>\t\n  <label for=\"product\">Product</label>\t\n  <input type=\"text\" id=\"product\" [(ngModel)]=\"agent.product\" name=\"product\">\n  <label for=\"expires\">Expiration Date</label>\t\n  <input type=\"date\" id=\"expires\" [(ngModel)]=\"agent.stateExpires\" name=\"stateExpires\">\n  <footer>\n    <button class=\"submit\" type=\"submit\">Add Agent</button>\n    <p *ngIf=\"success\"><i class=\"fa fa-check\"></i>Your agent is successfully registered</p>\n    <p *ngIf=\"failed\"><i class=\"fa fa-times\" aria-hidden=\"true\"></i>\n      Failed to register this agent</p>\n</footer>\n</form>\n</div>\n</div>"

/***/ }),

/***/ "../../../../../src/app/views/add-agent/add-agent.component.ts":
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "a", function() { return AddAgentComponent; });
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_0__angular_core__ = __webpack_require__("../../../core/esm5/core.js");
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_1__services_main_service__ = __webpack_require__("../../../../../src/app/services/main.service.ts");
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_2__models_agent__ = __webpack_require__("../../../../../src/app/models/agent.ts");
var __decorate = (this && this.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};
var __metadata = (this && this.__metadata) || function (k, v) {
    if (typeof Reflect === "object" && typeof Reflect.metadata === "function") return Reflect.metadata(k, v);
};



var AddAgentComponent = (function () {
    function AddAgentComponent(mainService) {
        this.mainService = mainService;
        this.agent = new __WEBPACK_IMPORTED_MODULE_2__models_agent__["a" /* Agent */]();
    }
    AddAgentComponent.prototype.onSubmit = function () {
        var _this = this;
        this.mainService.addAgent(this.agent)
            .subscribe(function (res) {
            console.log(res);
            if (res.status === 200) {
                console.log('Your agent is successfully registered');
                _this.success = true;
            }
        }, function (err) {
            console.log(err);
            _this.failed = true;
        });
    };
    AddAgentComponent = __decorate([
        Object(__WEBPACK_IMPORTED_MODULE_0__angular_core__["Component"])({
            selector: 'app-add-agent',
            template: __webpack_require__("../../../../../src/app/views/add-agent/add-agent.component.html"),
            styles: [__webpack_require__("../../../../../src/app/views/add-agent/add-agent.component.css")]
        }),
        __metadata("design:paramtypes", [__WEBPACK_IMPORTED_MODULE_1__services_main_service__["a" /* MainService */]])
    ], AddAgentComponent);
    return AddAgentComponent;
}());



/***/ }),

/***/ "../../../../../src/app/views/auth/auth.component.css":
/***/ (function(module, exports, __webpack_require__) {

exports = module.exports = __webpack_require__("../../../../css-loader/lib/css-base.js")(false);
// imports


// module
exports.push([module.i, ".wrapper {\r\n    height: 100vh;\r\n    width: 80%;\r\n    margin: 0 auto;\r\n    display: -webkit-box;\r\n    display: -ms-flexbox;\r\n    display: flex;\r\n    -webkit-box-pack: justify;\r\n        -ms-flex-pack: justify;\r\n            justify-content: space-between;\r\n    -webkit-box-align: center;\r\n        -ms-flex-align: center;\r\n            align-items: center;\r\n}\r\n\r\nimg {\r\n    -webkit-box-flex: 1;\r\n        -ms-flex: 1 1 auto;\r\n            flex: 1 1 auto;\r\n    max-width: 40%;\r\n}\r\n\r\n", ""]);

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
        Object(__WEBPACK_IMPORTED_MODULE_0__angular_core__["Component"])({
            selector: 'app-auth',
            template: __webpack_require__("../../../../../src/app/views/auth/auth.component.html"),
            styles: [__webpack_require__("../../../../../src/app/views/auth/auth.component.css")]
        }),
        __metadata("design:paramtypes", [])
    ], AuthComponent);
    return AuthComponent;
}());



/***/ }),

/***/ "../../../../../src/app/views/help/help.component.css":
/***/ (function(module, exports, __webpack_require__) {

exports = module.exports = __webpack_require__("../../../../css-loader/lib/css-base.js")(false);
// imports


// module
exports.push([module.i, ".content {\r\n    padding: 13vh 10vw;\r\n    display: -webkit-box;\r\n    display: -ms-flexbox;\r\n    display: flex;\r\n    -webkit-box-orient: horizontal;\r\n    -webkit-box-direction: normal;\r\n        -ms-flex-flow: row wrap;\r\n            flex-flow: row wrap;\r\n}\r\n\r\narticle {\r\n    width: 25%;\r\n    box-shadow: 0 0 5px rgb(23, 43, 43);\r\n    padding: 2rem;\r\n    background-color: white;\r\n    margin: 1.5rem;\r\n}\r\n", ""]);

// exports


/*** EXPORTS FROM exports-loader ***/
module.exports = module.exports.toString();

/***/ }),

/***/ "../../../../../src/app/views/help/help.component.html":
/***/ (function(module, exports) {

module.exports = "<app-header></app-header>\n<section class=\"content\">\n  <article *ngFor=\"let product of products\">\n    <app-item [product]=\"product\"></app-item>\n  </article>\n</section>\n<app-add-product-modal></app-add-product-modal>"

/***/ }),

/***/ "../../../../../src/app/views/help/help.component.ts":
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "a", function() { return HelpComponent; });
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


var HelpComponent = (function () {
    function HelpComponent(mainService) {
        this.mainService = mainService;
    }
    HelpComponent.prototype.ngOnInit = function () {
        var _this = this;
        this.mainService.getProductList().subscribe(function (data) {
            _this.products = data;
            console.log(data);
        });
    };
    HelpComponent = __decorate([
        Object(__WEBPACK_IMPORTED_MODULE_0__angular_core__["Component"])({
            selector: 'app-help',
            template: __webpack_require__("../../../../../src/app/views/help/help.component.html"),
            styles: [__webpack_require__("../../../../../src/app/views/help/help.component.css")]
        }),
        __metadata("design:paramtypes", [__WEBPACK_IMPORTED_MODULE_1__services_main_service__["a" /* MainService */]])
    ], HelpComponent);
    return HelpComponent;
}());



/***/ }),

/***/ "../../../../../src/app/views/home/home.component.css":
/***/ (function(module, exports, __webpack_require__) {

exports = module.exports = __webpack_require__("../../../../css-loader/lib/css-base.js")(false);
// imports


// module
exports.push([module.i, ".card {\r\n    width: 25%;\r\n    box-shadow: 0 0 5px rgb(23, 43, 43);\r\n    padding: 2rem 1rem;\r\n    background-color: white;\r\n    margin: 1.5rem;\r\n}\r\n\r\n.card > * {\r\n    margin: 1rem;\r\n}\r\n\r\n.content {\r\n    padding: 13vh 10vw;\r\n    display: -webkit-box;\r\n    display: -ms-flexbox;\r\n    display: flex;\r\n    -webkit-box-orient: horizontal;\r\n    -webkit-box-direction: normal;\r\n        -ms-flex-flow: row wrap;\r\n            flex-flow: row wrap;\r\n}\r\n\r\nheader {\r\n    display: -webkit-box;\r\n    display: -ms-flexbox;\r\n    display: flex;\r\n    -webkit-box-pack: justify;\r\n        -ms-flex-pack: justify;\r\n            justify-content: space-between;\r\n}\r\n\r\n.title {\r\n    -webkit-box-flex: 1;\r\n        -ms-flex: 1 1 60%;\r\n            flex: 1 1 60%;\r\n}\r\n\r\n.options {\r\n    display: -webkit-box;\r\n    display: -ms-flexbox;\r\n    display: flex;\r\n    -webkit-box-align: center;\r\n        -ms-flex-align: center;\r\n            align-items: center;\r\n    -webkit-box-flex: 1;\r\n        -ms-flex: 1 0 auto;\r\n            flex: 1 0 auto;\r\n    -webkit-box-pack: end;\r\n        -ms-flex-pack: end;\r\n            justify-content: flex-end;\r\n    font-size: 1.2rem;\r\n}\r\n\r\n.options > * {\r\n    cursor: pointer;\r\n}\r\n\r\n.options > *:hover {\r\n    color: #7ea1a1;\r\n    -webkit-transform: translateY(1px);\r\n            transform: translateY(1px);\r\n}\r\n\r\n.search {\r\n    margin-right: 0.5rem;\r\n}\r\n\r\n.search-accented {\r\n    text-shadow: 0 0 7px #008686;\r\n}\r\n\r\n.product-image {\r\n    width: 100%;\r\n}\r\n\r\n.product-expired {\r\n    border: 2px solid rgb(216, 45, 45);\r\n    -webkit-filter: grayscale(100%);\r\n            filter: grayscale(100%);\r\n}\r\n\r\n.text-ok {\r\n    color: #84d511;\r\n}\r\n\r\n.text-warn {\r\n    color:  rgb(233, 233, 33);\r\n}\r\n\r\n.text-alert {\r\n    color: rgb(216, 45, 45);\r\n}\r\n\r\n.noAgents {\r\n    display: -webkit-box;\r\n    display: -ms-flexbox;\r\n    display: flex;\r\n    -webkit-box-align: center;\r\n        -ms-flex-align: center;\r\n            align-items: center;\r\n    -webkit-box-pack: space-evenly;\r\n        -ms-flex-pack: space-evenly;\r\n            justify-content: space-evenly;\r\n    height: 52vh;\r\n    -webkit-box-orient: vertical;\r\n    -webkit-box-direction: normal;\r\n        -ms-flex-direction: column;\r\n            flex-direction: column;\r\n}\r\n\r\n.addAgent {\r\n    font-size: 1.1rem;\r\n    color: #fff;\r\n    border-radius: 30px;\r\n    outline: none;\r\n    border: 2px solid #fff;\r\n    width: 20%;\r\n    height: 8%;\r\n    padding: 1rem;\r\n    background-color: darkslategrey;\r\n    cursor: pointer;\r\n    text-align: center;\r\n}\r\n\r\n.addAgent:hover {\r\n    color: #cdd6d6;\r\n    background-color: #367373;\r\n}\r\n\r\neditModal {\r\n    position: fixed;\r\n    top: 50%;\r\n    left: 50%;\r\n    -webkit-transform: translate(-50%, -50%);\r\n            transform: translate(-50%, -50%);\r\n    background-color: white;\r\n    padding: 3rem;\r\n    box-shadow: 0 0 10px darkslategrey;\r\n}\r\n", ""]);

// exports


/*** EXPORTS FROM exports-loader ***/
module.exports = module.exports.toString();

/***/ }),

/***/ "../../../../../src/app/views/home/home.component.html":
/***/ (function(module, exports) {

module.exports = "<app-header></app-header>\r\n<section class=\"content\">\r\n    <div class=\"card\" *ngFor=\"let product of products\">\r\n    <app-product [product]=\"product\"\r\n    [ngClass]=\"{\r\n        'expiresSoon':product.condition === 'warn',\r\n        'expired':product.condition === 'expired'\r\n    }\">\r\n    </app-product>\r\n   </div>\r\n</section>\r\n\r\n<div class=\"noAgents\" *ngIf=\"noAgents\">\r\n    <h2>Seems like you don't have any agents registered</h2>\r\n    <div class=\"addAgent\" (click)=\"newAgentRedirect();\">Add New Agent</div>\r\n</div>\r\n"

/***/ }),

/***/ "../../../../../src/app/views/home/home.component.ts":
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "a", function() { return HomeComponent; });
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_0__angular_core__ = __webpack_require__("../../../core/esm5/core.js");
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_1__angular_router__ = __webpack_require__("../../../router/esm5/router.js");
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_2__services_main_service__ = __webpack_require__("../../../../../src/app/services/main.service.ts");
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
    function HomeComponent(mainService, router) {
        this.mainService = mainService;
        this.router = router;
        this.noAgents = false;
        this.success = false;
        this.failed = false;
    }
    HomeComponent.prototype.ngOnInit = function () {
        var _this = this;
        this.mainService.getProducts().subscribe(function (data) {
            _this.products = data;
            console.log(data);
            if (_this.products.length === 0) {
                _this.noAgents = true;
            }
        });
    };
    HomeComponent.prototype.newAgentRedirect = function () {
        this.router.navigate(['/newAgent']);
    };
    HomeComponent = __decorate([
        Object(__WEBPACK_IMPORTED_MODULE_0__angular_core__["Component"])({
            selector: 'app-home',
            template: __webpack_require__("../../../../../src/app/views/home/home.component.html"),
            styles: [__webpack_require__("../../../../../src/app/views/home/home.component.css")]
        }),
        __metadata("design:paramtypes", [__WEBPACK_IMPORTED_MODULE_2__services_main_service__["a" /* MainService */],
            __WEBPACK_IMPORTED_MODULE_1__angular_router__["Router"]])
    ], HomeComponent);
    return HomeComponent;
}());



/***/ }),

/***/ "../../../../../src/app/views/recipes/recipes.component.css":
/***/ (function(module, exports, __webpack_require__) {

exports = module.exports = __webpack_require__("../../../../css-loader/lib/css-base.js")(false);
// imports


// module
exports.push([module.i, ".card {\r\n    width: 450px;\r\n    box-shadow: 0 0 5px rgb(23, 43, 43);\r\n    padding: 2rem 1rem;\r\n    background-color: white;\r\n    margin: 2rem 1rem;\r\n}\r\n\r\n.card > * {\r\n    margin: 1rem;\r\n}\r\n\r\n.content {\r\n    padding: 13vh 10vw;\r\n    display: -webkit-box;\r\n    display: -ms-flexbox;\r\n    display: flex;\r\n    -webkit-box-orient: horizontal;\r\n    -webkit-box-direction: normal;\r\n        -ms-flex-flow: row wrap;\r\n            flex-flow: row wrap;\r\n    -webkit-box-pack: justify;\r\n        -ms-flex-pack: justify;\r\n            justify-content: space-between;\r\n}\r\n\r\nheader {\r\n    display: -webkit-box;\r\n    display: -ms-flexbox;\r\n    display: flex;\r\n    -webkit-box-pack: justify;\r\n        -ms-flex-pack: justify;\r\n            justify-content: space-between;\r\n}\r\n\r\n.title {\r\n    -webkit-box-flex: 1;\r\n        -ms-flex: 1 1 60%;\r\n            flex: 1 1 60%;\r\n}\r\n\r\n.tags {\r\n    display: -webkit-box;\r\n    display: -ms-flexbox;\r\n    display: flex;\r\n    -webkit-box-align: center;\r\n        -ms-flex-align: center;\r\n            align-items: center;\r\n    -webkit-box-flex: 1;\r\n        -ms-flex: 1 0 auto;\r\n            flex: 1 0 auto;\r\n    -webkit-box-pack: justify;\r\n        -ms-flex-pack: justify;\r\n            justify-content: space-between;\r\n}\r\n\r\n.complexity {\r\n    padding: 0.3rem 0.3rem;\r\n    border-radius: 15px;\r\n    -ms-flex-line-pack: center;\r\n        align-content: center;\r\n}\r\n\r\n.complexity-easy {\r\n    background-color: #84d511;\r\n}\r\n\r\n.complexity-normal {\r\n    background-color: rgb(255, 255, 43);\r\n    color: #4d4141;\r\n}\r\n\r\n.complexity-hard {\r\n    background-color: rgb(216, 45, 45);\r\n    color: #ffe2e2;\r\n}\r\n\r\n.description {\r\n    margin-bottom: 0.7rem;\r\n}", ""]);

// exports


/*** EXPORTS FROM exports-loader ***/
module.exports = module.exports.toString();

/***/ }),

/***/ "../../../../../src/app/views/recipes/recipes.component.html":
/***/ (function(module, exports) {

module.exports = "<app-header></app-header>\n<section class=\"content\">\n    <article class=\"card\" *ngFor=\"let recipe of recipes\">\n        <header>\n            <h3 class=\"title\">{{ recipe.title }}</h3>\n            <div class=\"tags\">\n                <div class=\"time\">\n                    <i class=\"fa fa-clock-o\" aria-hidden=\"true\"></i>\n                    {{ recipe.coockingTimeMin }} min</div>\n                <div class=\"complexity\" [ngClass]=\"{\n                    'complexity-easy':recipe.complexity === 'easy',\n                    'complexity-normal':recipe.complexity === 'normal',\n                    'complexity-hard':recipe.complexity === 'hard'\n                }\">\n                    <i class=\"fa\" [ngClass]=\"{\n                        'fa-check':recipe.complexity === 'easy',\n                        'fa-cog':recipe.complexity === 'normal',\n                        'fa-cogs':recipe.complexity === 'hard'\n                    }\" arria-hidden=\"true\"></i>\n                {{ recipe.complexity }}\n                </div>\n            </div>\n        </header>\n        <main>\n            <div class=\"description\">{{ recipe.description }}</div>\n            <div class=\"ingredients\">{{ recipe.ingredients }}</div>\n        </main>\n    </article>\n</section>"

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
        var _this = this;
        this.mainService.getRecipes().subscribe(function (data) {
            _this.recipes = data;
        });
    };
    RecipesComponent = __decorate([
        Object(__WEBPACK_IMPORTED_MODULE_0__angular_core__["Component"])({
            selector: 'app-recipes',
            template: __webpack_require__("../../../../../src/app/views/recipes/recipes.component.html"),
            styles: [__webpack_require__("../../../../../src/app/views/recipes/recipes.component.css")]
        }),
        __metadata("design:paramtypes", [__WEBPACK_IMPORTED_MODULE_1__services_main_service__["a" /* MainService */]])
    ], RecipesComponent);
    return RecipesComponent;
}());



/***/ }),

/***/ "../../../../../src/app/views/searchrecipes/searchrecipes.component.css":
/***/ (function(module, exports, __webpack_require__) {

exports = module.exports = __webpack_require__("../../../../css-loader/lib/css-base.js")(false);
// imports


// module
exports.push([module.i, ".card {\r\n    width: 450px;\r\n    box-shadow: 0 0 5px rgb(23, 43, 43);\r\n    padding: 2rem 1rem;\r\n    background-color: white;\r\n    margin: 2rem 1rem;\r\n}\r\n\r\n.card > * {\r\n    margin: 1rem;\r\n}\r\n\r\n.content {\r\n    padding: 13vh 10vw;\r\n    display: -webkit-box;\r\n    display: -ms-flexbox;\r\n    display: flex;\r\n    -webkit-box-orient: horizontal;\r\n    -webkit-box-direction: normal;\r\n        -ms-flex-flow: row wrap;\r\n            flex-flow: row wrap;\r\n    -webkit-box-pack: justify;\r\n        -ms-flex-pack: justify;\r\n            justify-content: space-between;\r\n}\r\n\r\nheader {\r\n    display: -webkit-box;\r\n    display: -ms-flexbox;\r\n    display: flex;\r\n    -webkit-box-pack: justify;\r\n        -ms-flex-pack: justify;\r\n            justify-content: space-between;\r\n}\r\n\r\n.title {\r\n    -webkit-box-flex: 1;\r\n        -ms-flex: 1 1 60%;\r\n            flex: 1 1 60%;\r\n}\r\n\r\n.tags {\r\n    display: -webkit-box;\r\n    display: -ms-flexbox;\r\n    display: flex;\r\n    -webkit-box-align: center;\r\n        -ms-flex-align: center;\r\n            align-items: center;\r\n    -webkit-box-flex: 1;\r\n        -ms-flex: 1 0 auto;\r\n            flex: 1 0 auto;\r\n    -webkit-box-pack: justify;\r\n        -ms-flex-pack: justify;\r\n            justify-content: space-between;\r\n}\r\n\r\n.complexity {\r\n    padding: 0.3rem 0.3rem;\r\n    border-radius: 15px;\r\n    -ms-flex-line-pack: center;\r\n        align-content: center;\r\n}\r\n\r\n.complexity-easy {\r\n    background-color: #84d511;\r\n}\r\n\r\n.complexity-normal {\r\n    background-color: rgb(255, 255, 43);\r\n    color: #4d4141;\r\n}\r\n\r\n.complexity-hard {\r\n    background-color: rgb(216, 45, 45);\r\n    color: #ffe2e2;\r\n}\r\n\r\n.description {\r\n    margin-bottom: 0.7rem;\r\n}", ""]);

// exports


/*** EXPORTS FROM exports-loader ***/
module.exports = module.exports.toString();

/***/ }),

/***/ "../../../../../src/app/views/searchrecipes/searchrecipes.component.html":
/***/ (function(module, exports) {

module.exports = "<app-header></app-header>\n<app-header></app-header>\n<section class=\"content\">\n    <article class=\"card\" *ngFor=\"let recipe of recipes\">\n        <header>\n            <h3 class=\"title\">{{ recipe.title }}</h3>\n            <div class=\"tags\">\n                <div class=\"time\">\n                    <i class=\"fa fa-clock-o\" aria-hidden=\"true\"></i>\n                    {{ recipe.coockingTimeMin }} min</div>\n                <div class=\"complexity\" [ngClass]=\"{\n                    'complexity-easy':recipe.complexity === 'easy',\n                    'complexity-normal':recipe.complexity === 'normal',\n                    'complexity-hard':recipe.complexity === 'hard'\n                }\">\n                    <i class=\"fa\" [ngClass]=\"{\n                        'fa-check':recipe.complexity === 'easy',\n                        'fa-cog':recipe.complexity === 'normal',\n                        'fa-cogs':recipe.complexity === 'hard'\n                    }\" arria-hidden=\"true\"></i>\n                {{ recipe.complexity }}\n                </div>\n            </div>\n        </header>\n        <main>\n            <div class=\"description\">{{ recipe.description }}</div>\n            <div class=\"ingredients\">{{ recipe.ingredients }}</div>\n        </main>\n    </article>\n</section>"

/***/ }),

/***/ "../../../../../src/app/views/searchrecipes/searchrecipes.component.ts":
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "a", function() { return SearchrecipesComponent; });
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


var SearchrecipesComponent = (function () {
    function SearchrecipesComponent(mainService) {
        this.mainService = mainService;
    }
    SearchrecipesComponent.prototype.ngOnInit = function () {
        var _this = this;
        this.mainService.getMyRecipes().subscribe(function (data) {
            _this.recipes = data;
        });
    };
    SearchrecipesComponent = __decorate([
        Object(__WEBPACK_IMPORTED_MODULE_0__angular_core__["Component"])({
            selector: 'app-searchrecipes',
            template: __webpack_require__("../../../../../src/app/views/searchrecipes/searchrecipes.component.html"),
            styles: [__webpack_require__("../../../../../src/app/views/searchrecipes/searchrecipes.component.css")]
        }),
        __metadata("design:paramtypes", [__WEBPACK_IMPORTED_MODULE_1__services_main_service__["a" /* MainService */]])
    ], SearchrecipesComponent);
    return SearchrecipesComponent;
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
    Object(__WEBPACK_IMPORTED_MODULE_0__angular_core__["enableProdMode"])();
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