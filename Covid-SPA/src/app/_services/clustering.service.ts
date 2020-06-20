import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { environment } from 'src/environments/environment';
import { GroupedData } from '../_models/grouped-data';

@Injectable({
  providedIn: 'root'
})
export class ClusteringService {

  baseUrl = environment.apiUrl;

  constructor(private http: HttpClient) { }

  getClusters(k: number): Observable<GroupedData> {
    return this.http.get<GroupedData>(this.baseUrl + 'clusters/' + k);
  }

}