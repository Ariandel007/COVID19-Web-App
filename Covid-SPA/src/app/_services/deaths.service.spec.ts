
import { TestBed, async, inject } from '@angular/core/testing';
import { AnalisisService } from './analisis.service';

describe('Service: Diagnostic', () => {
  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [AnalisisService]
    });
  });

  it('should ...', inject([AnalisisService], (service: AnalisisService) => {
    expect(service).toBeTruthy();
  }));
});
