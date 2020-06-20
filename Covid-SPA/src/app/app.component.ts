import { Component } from '@angular/core';
import { ListDeaths } from './_models/list-deaths';
import { GroupedData } from './_models/grouped-data';
import { ClusteringService } from './_services/clustering.service';
import { DeathsService } from './_services/deaths.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'Covid-SPA';

  deaths: ListDeaths;
  groupedData?: GroupedData;
  k: number;
  seccion = 1;


  constructor(private clusteringService: ClusteringService, private deathsService: DeathsService) {}

  ngOnInit(): void {
    this.obtenerListaDeaths();
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

  obtenerListaDeaths(): void {
    this.deathsService.getDeaths().subscribe( (response) => {
      this.deaths = response;
    }, error => {
      console.log(error);
    });
  }

  setSeccion(n: number) {
    this.seccion = n;
  }
}
