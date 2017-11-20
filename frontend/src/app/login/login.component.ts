import { Component, OnInit, ViewEncapsulation } from '@angular/core';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss']
})
export class LoginComponent implements OnInit {

    button_text: String = "Giriş";

  constructor() { }

  ngOnInit() {
  }

  onSubmit(username: String) {
      console.log("login", username);
      this.button_text = "Giriş Yapılıyor: " + username;
  }

}
