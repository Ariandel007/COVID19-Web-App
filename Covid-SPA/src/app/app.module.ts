import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { HttpClientModule } from '@angular/common/http';
import {FormsModule} from '@angular/forms';


import { AppComponent } from './app.component';
import { ClusteringService } from './_services/clustering.service';
import { AnalisisService } from './_services/analisis.service';

@NgModule({
  declarations: [
    AppComponent
  ],
  imports: [
    BrowserModule,
    HttpClientModule,
    FormsModule
  ],
  providers: [
    ClusteringService,
    AnalisisService
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
