import { TestBed, async, inject } from '@angular/core/testing';
import { ClusteringService } from './clustering.service';

describe('Service: Clustering', () => {
  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [ClusteringService]
    });
  });

  it('should ...', inject([ClusteringService], (service: ClusteringService) => {
    expect(service).toBeTruthy();
  }));
});