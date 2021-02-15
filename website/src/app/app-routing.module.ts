import { RouterModule, Routes } from '@angular/router';
import { NgModule } from '@angular/core';

import { RetrievePasswordComponent } from './pages/retrieve-password/retrieve-password.component';
import { RegisterComponent } from './pages/register/register.component';
import { HomePageComponent } from '@pages/home-page/home-page.component';
import { NotFoundComponent } from '@pages/not-found/not-found.component';
import { LoginComponent } from '@pages/login/login.component';

const routes: Routes = [
  { path: "", component: HomePageComponent },
  { path: "retrieve-password", component: RetrievePasswordComponent },
  { path: "register", component: RegisterComponent },
  { path: "login", component: LoginComponent },
  { path: "**", component: NotFoundComponent },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
