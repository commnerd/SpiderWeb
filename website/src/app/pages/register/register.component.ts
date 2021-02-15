import { FormControl, FormGroup } from '@angular/forms';
import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.scss']
})
export class RegisterComponent implements OnInit {

  registerForm : FormGroup = new FormGroup({
    'email': new FormControl(""),
    'password': new FormControl(""),
  });
  
  constructor() { }

  ngOnInit(): void {
  }

}
