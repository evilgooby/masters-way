import clsx from "clsx";
import {languageStore} from "src/globalStore/LanguageStore";
import {LanguageService} from "src/service/LanguageService";
import {convertMinutesToHours} from "src/utils/convertMinutesToHours";
import styles from "src/logic/wayPage/wayStatistics/statisticBlock/statisticItem/StatisticItem.module.scss";

/**
 * Type of statisticItem's styles
 */
export enum StatisticItemType {

  /**
   * Primary statisticItem
   */
  PRIMARY = "primary",

  /**
   * Secondary statisticItem
   */
  SECONDARY = "secondary"
}

/**
 * StatisticItem
 */
export interface StatisticItem {

  /**
   * StatisticItem value
   */
  value: number;

  /**
   * StatisticItem text
   */
  text: string;

}

/**
 *StatisticItem props
 */
interface StatisticItemProps {

  /**
   * StatisticItem
   */
  statisticItem: StatisticItem;

  /**
   * StatisticItem className
   */
  type?: string;

  /**
   * Unit to which minutes will be converted
   */
  convertToHours?: boolean;
}

/**
 * StatisticItem
 */
export const StatisticItem = (props: StatisticItemProps) => {
  const {language} = languageStore;

  /**
   * Get formatted statistic value
   */
  const getFormattedStatisticValue = (): string => {
    if (props.type === StatisticItemType.SECONDARY || props.convertToHours) {
      const convertedMinutes = convertMinutesToHours(props.statisticItem.value);

      return `${convertedMinutes}${LanguageService.way.statisticsBlock.unitOfMeasurement[language]}`;
    }

    return props.statisticItem.value.toString();
  };

  return (
    <div className={clsx(styles.statisticItem, styles[props.type ?? StatisticItemType.PRIMARY])}>
      <span className={styles.statisticValue}>
        {getFormattedStatisticValue()}
      </span>
      <span className={styles.statisticText}>
        {props.statisticItem.text}
      </span>
    </div>
  );
};
