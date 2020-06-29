import { Component, OnInit } from '@angular/core';
import { ListAnalisis } from './_models/list-analisis';
import { GroupedData } from './_models/grouped-data';
import { ClusteringService } from './_services/clustering.service';
import { AnalisisService } from './_services/analisis.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements OnInit  {
  title = 'Covid-SPA';

  analisis: ListAnalisis;
  groupedData?: GroupedData;
  k: number;
  seccion = 1;

  //KNN
  prediccion:number;
  kneighbors:any;

  paciente:any;

  constructor(private clusteringService: ClusteringService, private analisisService: AnalisisService) {}

  ngOnInit(): void {
    this.obtenerListaAnalisis();
  }

  agruparDatos(): void{
    console.log(this.k);
    if ( this.k <= 0 ) {
      return;
    }

    this.clusteringService.getClusters(this.k).subscribe( (response) => {
      this.groupedData = response;
    }, error => {
      console.log(error);
    });
  }

  obtenerListaAnalisis(): void {
    this.analisisService.getAnalisis().subscribe( (response) => {
      this.analisis = response;
    }, error => {
      console.log(error);
    });
  }

  setSeccion(n: number) {
    this.seccion = n;
  }

  realizarPrediccion():void{
    console.log(this.kneighbors);
     this.analisisService.getPrediccion(this.kneighbors,this.paciente).subscribe((response)=>{
       console.log(response)
       this.prediccion=response.prediccion
     })
  }
}
