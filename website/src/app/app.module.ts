import { BrowserModule } from '@angular/platform-browser';
import { ReactiveFormsModule } from '@angular/forms';
import { NgModule } from '@angular/core';

import { HomePageComponent } from './pages/home-page/home-page.component';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { NotFoundComponent } from './pages/not-found/not-found.component';
import { NavBarComponent } from './partials/nav-bar/nav-bar.component';
import { HomeComponent } from './layouts/home/home.component';
import { LoginComponent } from './pages/login/login.component';
import { StandardComponent } from './layouts/standard/standard.component';
import { LinkComponent } from './partials/link/link.component';
import { RegisterComponent } from './pages/register/register.component';
import { RetrievePasswordComponent } from './pages/retrieve-password/retrieve-password.component';

@NgModule({
  declarations: [
    AppComponent,
    HomePageComponent,
    NotFoundComponent,
    LinkComponent,
    NavBarComponent,
    HomeComponent,
    LoginComponent,
    StandardComponent,
    RegisterComponent,
    RetrievePasswordComponent,
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    ReactiveFormsModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
