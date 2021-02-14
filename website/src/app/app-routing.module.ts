import { RouterModule, Routes } from '@angular/router';
import { HomePageComponent } from '@pages/home-page/home-page.component';
import { NotFoundComponent } from '@pages/not-found/not-found.component';
import { NgModule } from '@angular/core';

const routes: Routes = [
  { path: "", component: HomePageComponent },
  { path: "**", component: NotFoundComponent },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
