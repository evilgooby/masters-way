import {SchemasDefaultWayCollections} from "src/apiAutogenerated";
import {wayCollectionDTOToWayCollection} from "src/dataAccessLogic/DTOToPreviewConverter/wayCollectionDTOToWayCollection";
import {DefaultWayCollections} from "src/model/businessModel/User";

/**
 * Convert {@link SchemasDefaultWayCollections} to {@link DefaultWayCollections}
 */
export const defaultWayCollectionsDTOToDefaultWayCollections = (
  defaultWayCollections: SchemasDefaultWayCollections,
): DefaultWayCollections => {
  return new DefaultWayCollections({
    own: wayCollectionDTOToWayCollection(defaultWayCollections.own),
    favorite: wayCollectionDTOToWayCollection(defaultWayCollections.favorite),
    mentoring: wayCollectionDTOToWayCollection(defaultWayCollections.mentoring),
  });
};
