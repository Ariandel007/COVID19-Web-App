
import { TestBed, async, inject } from '@angular/core/testing';
import { DeathsService } from './deaths.service';

describe('Service: Diagnostic', () => {
  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [DeathsService]
    });
  });

  it('should ...', inject([DeathsService], (service: DeathsService) => {
    expect(service).toBeTruthy();
  }));
});