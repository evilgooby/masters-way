import {makeAutoObservable} from "mobx";
import {WayPreview} from "src/model/businessModelPreview/WayPreview";
import {PartialWithUuid} from "src/utils/PartialWithUuid";

/**
 * WayCollection props
 */
interface WayCollectionProps {

  /**
   * Date in ISO format
   */
  createdAt: Date;

  /**
   * Way collection name
   */
  name: string;

  /**
   * Way collection's owner
   */
  ownerUuid: string;

  /**
   * Date in ISO format
   */
  updatedAt: Date;

  /**
   * Way collection's UUID
   */
  uuid: string;

  /**
   * Ways preview that exist inside way collection
   */
  ways: WayPreview[];

}

/**
 * Default WayCollections
 */
export class DefaultWayCollections {

  /**
   * Own collection
   */
  public own: WayCollection;

  /**
   * Favorite collection
   */
  public favorite: WayCollection;

  /**
   * Mentoring collection
   */
  public mentoring: WayCollection;

  constructor(wayCollections: DefaultWayCollections) {
    makeAutoObservable(this);
    this.own = wayCollections.own;
    this.favorite = wayCollections.favorite;
    this.mentoring = wayCollections.mentoring;

  }

}

/**
 * Specific way collection data
 */
export class WayCollection {

  /**
   * Date in ISO format
   */
  public createdAt: Date;

  /**
   * Way collection name
   */
  public name: string;

  /**
   * Way collection's owner
   */
  public ownerUuid: string;

  /**
   * Date in ISO format
   */
  public updatedAt: Date;

  /**
   * Way collection's UUID
   */
  public uuid: string;

  /**
   * Ways preview that exist inside way collection
   */
  public ways: WayPreview[];

  constructor(wayCollection: WayCollectionProps) {
    makeAutoObservable(this);
    this.createdAt = wayCollection.createdAt;
    this.name = wayCollection.name;
    this.ownerUuid = wayCollection.ownerUuid;
    this.updatedAt = wayCollection.updatedAt;
    this.uuid = wayCollection.uuid;
    this.ways = wayCollection.ways;
  }

  /**
   * Add way to ways
   */
  public addWay(newWay: WayPreview) {
    this.ways.push(newWay);
  }

  /**
   * Delete way to ways
   */
  public deleteWay(wayUuid: string) {
    const updatedWays = this.ways.filter((way) => way.uuid !== wayUuid);
    this.ways = updatedWays;
  }

  /**
   * Update way to ways
   */
  public updateWay(wayToUpdate: PartialWithUuid<WayPreview>) {
    const updatedWays = this.ways.map((way) => way.uuid !== wayToUpdate.uuid
      ? way
      : new WayPreview({...way, ...wayToUpdate}),
    );
    this.ways = updatedWays;
  }

  /**
   * Update way collection name
   */
  public updateWayCollectionName(name: string) {
    this.name = name;
  }

}

/**
 * User tag data
 */
export class UserTag {

  /**
   * User tag uuid
   */
  public uuid: string;

  /**
   * User tag name
   */
  public name: string;

  constructor(userTag: UserTag) {
    makeAutoObservable(this);
    this.name = userTag.name;
    this.uuid = userTag.uuid;
  }

}

/**
 * Main short user information
 */
export class UserPlain {

  /**
   * User's UUID
   */
  public uuid: string;

  /**
   * Date when user was created
   */
  public createdAt: Date;

  /**
   * User's description
   */
  public description: string;

  /**
   * User's e-mail
   */
  public email: string;

  /**
   * User's image path
   */
  public imageUrl: string | null;

  /**
   * Is user mentor or not
   */
  public isMentor: boolean;

  /**
   * User's name
   */
  public name: string;

  constructor(user: UserPlain) {
    makeAutoObservable(this);
    this.uuid = user.uuid;
    this.createdAt = user.createdAt;
    this.description = user.description;
    this.email = user.email;
    this.imageUrl = user.imageUrl;
    this.isMentor = user.isMentor;
    this.name = user.name;
  }

}

/**
 * User props
 */
interface UserProps {

  /**
   * User's UUID
   */
  uuid: string;

  /**
   * User's name
   */
  name: string;

  /**
   * User's e-mail
   */
  email: string;

  /**
   * User's description
   */
  description: string;

  /**
   * Date when user was created
   */
  createdAt: Date;

  /**
   * Custom way collections {@link WayCollection}
   */
  customWayCollections: WayCollection[];

  /**
   * Default way collections {@link DefaultWayCollections}
   */
  defaultWayCollections: DefaultWayCollections;

  /**
   * User's uuids for whom this user are favorite
   */
  favoriteForUserUuids: string[];

  /**
   * Uuids of users who you liked
   */
  favoriteUsers: UserPlain[];

  /**
   * User's tags {@link UserTag}
   */
  tags: UserTag[];

  /**
   * User's image path
   */
  imageUrl: string;

  /**
   * Is user mentor or not
   */
  isMentor: boolean;

  /**
   * Way's that requested user become a mentor {@link WayPreview}
   */
  wayRequests: WayPreview[];
}

/**
 * User preview model
 */
export class User {

  /**
   * User's UUID
   */
  public uuid: string;

  /**
   * User's name
   */
  public name: string;

  /**
   * User's e-mail
   */
  public email: string;

  /**
   * User's description
   */
  public description: string;

  /**
   * Date when user was created
   */
  public createdAt: Date;

  /**
   * Custom way collections {@link WayCollection}
   */
  public customWayCollections: WayCollection[];

  /**
   * Default way collections {@link DefaultWayCollections}
   */
  public defaultWayCollections: DefaultWayCollections;

  /**
   * User's uuids for whom this user are favorite
   */
  public favoriteForUserUuids: string[];

  /**
   * Uuids of users who you liked
   */
  public favoriteUsers: UserPlain[];

  /**
   * User's tags {@link UserTag}
   */
  public tags: UserTag[];

  /**
   * User's image path
   */
  public imageUrl: string;

  /**
   * Is user mentor or not
   */
  public isMentor: boolean;

  /**
   * Way's that requested user become a mentor {@link WayPreview}
   */
  public wayRequests: WayPreview[];

  constructor(userData: UserProps) {
    makeAutoObservable(this);
    this.uuid = userData.uuid;
    this.name = userData.name;
    this.email = userData.email;
    this.description = userData.description;
    this.createdAt = userData.createdAt;
    this.customWayCollections = userData.customWayCollections.map(customWayCollection => new WayCollection(customWayCollection));
    this.defaultWayCollections = new DefaultWayCollections(userData.defaultWayCollections);
    this.favoriteForUserUuids = userData.favoriteForUserUuids;
    this.favoriteUsers = userData.favoriteUsers;
    this.tags = userData.tags.map(tag => new UserTag(tag));
    this.imageUrl = userData.imageUrl;
    this.isMentor = userData.isMentor;
    this.wayRequests = userData.wayRequests;
  }

  /**
   * Update user's name
   */
  public updateName(nameToUpdate: string): void {
    this.name = nameToUpdate;
    this.updateDefaultWayCollections("name", nameToUpdate);
  }

  /**
   * Update isMentor property
   */
  public updateIsMentor(isMentorToUpdate: boolean): void {
    this.isMentor = isMentorToUpdate;
    this.updateDefaultWayCollections("isMentor", isMentorToUpdate);
  }

  /**
   * Update user's description
   */
  public updateDescription(descriptionToUpdate: string): void {
    this.description = descriptionToUpdate;
  }

  /**
   * Add new collection to user
   */
  public addCollection(newCollection: WayCollection): void {
    this.customWayCollections.push(newCollection);
  }

  /**
   * Delete collection from user
   */
  public deleteCollection(collectionId: string): void {
    this.customWayCollections = this.customWayCollections.filter(collection => collection.uuid !== collectionId);
  }

  /**
   * Add new skill to user
   */
  public addTag(newTag: UserTag): void {
    this.tags.push(newTag);
  }

  /**
   * Delete skill from user
   */
  public deleteTag(skillUuid: string): void {
    this.tags = this.tags.filter(tag => tag.uuid !== skillUuid);
  }

  /**
   * Add way to favorite collection
   */
  public addWayToFavorite(newWay: WayPreview): void {
    this.defaultWayCollections.favorite.addWay(newWay);
  }

  /**
   * Delete way from favorite collection
   */
  public deleteWayFromFavorite(wayUuid: string): void {
    this.defaultWayCollections.favorite.deleteWay(wayUuid);
  }

  /**
   * Add way to composite
   */
  public addWayToComposite(parentWayUuid: string, childWayUuid: string): void {
    const ways = this.defaultWayCollections.own.ways.map((way) => {

      return way.uuid !== parentWayUuid
        ? way
        : new WayPreview({
          ...way,
          childrenUuids: way.childrenUuids.concat(childWayUuid),
        });
    });
    const ownCollection = new WayCollection({...this.defaultWayCollections.own, ways});

    this.defaultWayCollections = new DefaultWayCollections({
      ...this.defaultWayCollections,
      own: ownCollection,
    });
  }

  /**
   * Delete way from composite
   */
  public deleteWayFromComposite(parentWayUuid: string, childWayUuid: string): void {
    const ways = this.defaultWayCollections.own.ways.map((way) => {

      return way.uuid !== parentWayUuid
        ? way
        : new WayPreview({
          ...way,
          childrenUuids: way.childrenUuids.filter(wayUuid => wayUuid !== childWayUuid),
        });
    });
    const ownCollection = new WayCollection({...this.defaultWayCollections.own, ways});

    this.defaultWayCollections = new DefaultWayCollections({
      ...this.defaultWayCollections,
      own: ownCollection,
    });
  }

  /**
   * Update DefaultWayCollections
   */
  private updateDefaultWayCollections(updateKey: keyof User, updateValue: User["name"] | User["isMentor"]) {
    const ways = this.defaultWayCollections.own.ways.map((way) => {
      const owner = new UserPlain({...way.owner, [updateKey]: updateValue});

      return new WayPreview({
        ...way,
        owner,
      });
    });
    const ownCollection = new WayCollection({...this.defaultWayCollections.own, ways});

    this.defaultWayCollections = new DefaultWayCollections({
      ...this.defaultWayCollections,
      own: ownCollection,
    });
  }

}
