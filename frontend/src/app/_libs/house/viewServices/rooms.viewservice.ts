import { computed, Injectable, signal, Signal, WritableSignal } from '@angular/core';
import Room from '../models/room.model';
import { RoomsService } from '../backend/services/rooms.service';
import { ApiResponse } from '../../backend/models/api-response.model';
import { SelectItemGroup } from 'primeng/api';
import { LoadingStatus } from '../enums/loading-status.enum';

@Injectable({
    providedIn: 'root',
})
export class RoomsViewService {
    public rooms: WritableSignal<Room[]> = signal<Room[]>([]);
    public loadingStatus: WritableSignal<LoadingStatus> = signal(LoadingStatus.NONE);

    public groupedRooms = computed(() => {
        const grouped: { [key: string]: Room[] } = {
            Garage: [],
            EG: [],
            OG: [],
            UG: [],
            Sonstige: [],
        };

        this.rooms().forEach((room) => {
            if (room.name === 'Garage') {
                grouped['Garage'].push(room);
            } else if (room.name.startsWith('EG')) {
                grouped['EG'].push(room);
            } else if (room.name.startsWith('OG')) {
                grouped['OG'].push(room);
            } else if (room.name.startsWith('UG')) {
                grouped['UG'].push(room);
            } else {
                grouped['Sonstige'].push(room);
            }
        });

        return grouped;
    });

    public groupedRoomsForDropdown = computed<SelectItemGroup[]>(() => {
        return Object.entries(this.groupedRooms())
            .filter(([_, rooms]) => rooms.length > 0)
            .map(([label, rooms]) => ({
                label,
                items: rooms.map((room) => ({ label: room.name, value: room })),
            }));
    });

    constructor(private _roomsService: RoomsService) {
        this.loadRooms();
    }

    private loadRooms() {
        this.loadingStatus.set(LoadingStatus.LOADING);
        this._roomsService.all().subscribe({
            next: (apiResponse: ApiResponse<Room[]>) => {
                if (apiResponse.isSuccessful && apiResponse.data) {
                    this.rooms.set(apiResponse.data);
                }
            },
            error: (error: any) => {
                // @todo: error handling
                this.loadingStatus.set(LoadingStatus.LOADING_ERROR);
                console.log(`Error loading rooms: ${error}`);
            },
            complete: () => {
                this.loadingStatus.set(LoadingStatus.LOADING_SUCCESS);
            },
        });
    }

    // constructor(private _roomsService: RoomsService) {
    //     this.rooms = undefined;

    //     this._roomsService.all().subscribe((apiResonse: ApiResponse<Room[]>) => {
    //         if (apiResonse.isSuccessful) {
    //             this.rooms = apiResonse.data;

    //             this.rooms?.forEach((room) => {
    //                 if (room.name === 'Garage') {
    //                     this.groupedRooms['Garage'].push(room);
    //                 }
    //                 if (room.name === 'OG') {
    //                     this.groupedRooms['OG'].push(room);
    //                 }
    //                 if (room.name === 'EG') {
    //                     this.groupedRooms['EG'].push(room);
    //                 }
    //                 if (room.name === 'UG') {
    //                     this.groupedRooms['UG'].push(room);
    //                 }
    //             });
    //         }
    //     });
    // }
}
