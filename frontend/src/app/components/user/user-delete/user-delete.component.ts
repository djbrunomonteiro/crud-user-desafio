import { UserService } from "./../user.service";
import { Component, OnInit } from "@angular/core";
import { User } from "../user.model";
import { Route } from "@angular/compiler/src/core";
import { ActivatedRoute, Router } from "@angular/router";

@Component({
  selector: "app-user-delete",
  templateUrl: "./user-delete.component.html",
  styleUrls: ["./user-delete.component.css"],
})
export class UserDeleteComponent implements OnInit {
  user!: User;
  displayedColumns = ['nome', 'email'];

  constructor(private router: Router, private userService: UserService, private route: ActivatedRoute) {}

  ngOnInit(): void {
    const _id= this.route.snapshot.paramMap.get('_id') ;

    this.userService.readById(_id!).subscribe(user =>{
      this.user = user
    })
  }

  deleteUser(): void{
    this.userService.delete(`${this.user._id}`).subscribe(()=>{
      this.userService.showMessage('Produto excluido com sucesso!');
      this.router.navigate(["/usuarios"]);

    })
  }
  cancel(): void {
    this.router.navigate(["/usuarios"]);
  }

}

