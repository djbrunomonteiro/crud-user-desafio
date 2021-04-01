import { User } from "./../user.model";
import { UserService } from "./../user.service";
import { Component, OnInit } from "@angular/core";
import { Router } from "@angular/router";

@Component({
  selector: "app-user-create",
  templateUrl: "./user-create.component.html",
  styleUrls: ["./user-create.component.css"],
})
export class UserCreateComponent implements OnInit {
  user: User = {
    _id: "",    
    nome: "",
    email: "",
    senha: ""
  };

  constructor(private userService: UserService, private router: Router) {}

  ngOnInit(): void {}

  createUser(): void {
    this.userService.create(this.user).subscribe(() => {
      this.userService.showMessage("Usuario criado!");
      this.router.navigate(["/usuarios"]);
    });
  }
  cancel(): void {
    this.router.navigate(["/usuarios"]);
  }
}
