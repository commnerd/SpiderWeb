import { FormControl, FormGroup } from '@angular/forms';
import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-retrieve-password',
  templateUrl: './retrieve-password.component.html',
  styleUrls: ['./retrieve-password.component.scss']
})
export class RetrievePasswordComponent implements OnInit {

  retrievePasswordForm: FormGroup = new FormGroup({
    'email': new FormControl(""),
  });

  constructor() { }

  ngOnInit(): void {
  }

}
