import { User } from './user.model';
import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { MatSnackBar} from '@angular/material/snack-bar';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class UserService {

  baseUrl = "http://localhost:7000/usuarios";

  constructor(private snackBar: MatSnackBar, private http: HttpClient) { }

  showMessage(msg: string): void {
    this.snackBar.open(msg, 'x',{
      duration: 3000, 
      horizontalPosition: "right", 
      verticalPosition: "top"
    })
  }

 
  create(user: User): Observable<User>{
    return this.http.post<User>(this.baseUrl, user)
  }

  read(): Observable<User[]>{
    
    return this.http.get<User[]>(this.baseUrl);
    

  }

  readById(_id: string): Observable<User> {
    const url = `${this.baseUrl}/delete/${_id}`
    return this.http.get<User>(url);
  }

  update(user: User): Observable<User> {
    const url = `${this.baseUrl}/${user._id}`
    return this.http.put<User>(url, user);
  }

  delete(_id: string): Observable<User> {
    const url = `${this.baseUrl}/${_id}`;
    return this.http.delete<User>(url);

  }
}
