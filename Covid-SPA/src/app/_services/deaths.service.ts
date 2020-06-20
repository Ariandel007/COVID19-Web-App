import { Injectable } from '@angular/core';
import { environment } from 'src/environments/environment';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs/internal/Observable';
import { ListDeaths } from '../_models/list-deaths';

@Injectable({
  providedIn: 'root'
})
export class DeathsService {


  baseUrl = environment.apiUrl;

  constructor(private http: HttpClient) { }

  getDeaths(): Observable<ListDeaths> {
    return this.http.get<ListDeaths>(this.baseUrl + 'deaths');
  }
}