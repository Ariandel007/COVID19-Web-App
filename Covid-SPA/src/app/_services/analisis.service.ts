import { Injectable } from '@angular/core';
import { environment } from 'src/environments/environment';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs/internal/Observable';
import { ListAnalisis } from '../_models/list-analisis';
import { Prediccion } from '../_models/prediccion';
import { Paciente } from '../_models/paciente';
@Injectable({
  providedIn: 'root'
})
export class AnalisisService {

  baseUrl = environment.apiUrl;

  constructor(private http: HttpClient) { }

  getAnalisis(): Observable<ListAnalisis> {
    return this.http.get<ListAnalisis>(this.baseUrl + 'data');
  }

  getPrediccion(k: number, paciente: Paciente) {
    return this.http.post(this.baseUrl + 'prediccion/' + k, paciente);
  }

  getDiagnosis() {
    return this.http.get(this.baseUrl + 'pred');
  }
}
