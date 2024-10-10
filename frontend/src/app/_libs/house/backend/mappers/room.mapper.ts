import { IconDefinition } from '@fortawesome/free-solid-svg-icons';
import Room from '../../models/room.model';
import { RoomDto } from '../dtos/room.dto';
import { RoomNames } from '../../enums/rooms.enum';
import { IconDataprovider } from '../../../icons/icon.dataprovider';

export function mapDtoToModelArray(dtos: RoomDto[] | undefined): Room[] | undefined {
    if (!dtos) {
        return undefined;
    }

    const rooms: Room[] = [];

    dtos.forEach((room) => {
        rooms.push(mapDtoToModel(room)!);
    });

    return rooms;
}

export function mapDtoToModel(dto: RoomDto | undefined): Room | undefined {
    if (!dto) {
        return undefined;
    }

    let icon: IconDefinition;

    switch (dto.name) {
        case RoomNames.GARAGE:
            icon = IconDataprovider.outside;
            break;
        case (RoomNames.GROUNDFLOOR_KITCHEN, RoomNames.UPPERFLOOR_KITCHEN):
            icon = IconDataprovider.kitchen;
            break;
        case (RoomNames.GROUNDFLOOR_LIVINGROOM, RoomNames.UPPERFLOOR_LIVINGROOM):
            icon = IconDataprovider.livingRoom;
            break;
        case RoomNames.UPPERFLOOR_BEDROOM:
            icon = IconDataprovider.bedroom;
            break;
        case RoomNames.UPPERFLOOR_GALLERY:
            icon = IconDataprovider.gallery;
            break;
        case RoomNames.BASEMENTFLOOR_GYM:
            icon = IconDataprovider.gym;
            break;
        default:
            icon = IconDataprovider.unknown;
    }

    return new Room(dto.id, dto.externalId, dto.name, icon);
}
