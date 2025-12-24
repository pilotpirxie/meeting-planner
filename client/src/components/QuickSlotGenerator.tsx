import { useState } from "react";
import type { TimeSlot } from "../types";

interface QuickSlotGeneratorProps {
  startDate: string;
  endDate: string;
  dailyStartTime: string;
  dailyEndTime: string;
  duration: string;
  isOverlapping: boolean;
  isWholeDay: boolean;
  onStartDateChange: (date: string) => void;
  onEndDateChange: (date: string) => void;
  onDailyStartTimeChange: (time: string) => void;
  onDailyEndTimeChange: (time: string) => void;
  onDurationChange: (duration: string) => void;
  onOverlappingChange: (isOverlapping: boolean) => void;
  onWholeDayChange: (isWholeDay: boolean) => void;
}

export const useQuickSlotGenerator = () => {
  const [isQuickModalOpen, setIsQuickModalOpen] = useState(false);
  const [quickData, setQuickData] = useState({
    startDate: "",
    endDate: "",
    dailyStartTime: "",
    dailyEndTime: "",
    duration: "",
    isOverlapping: true,
    isWholeDay: false,
  });

  const handleOpenQuickModal = () => {
    setQuickData({
      startDate: "",
      endDate: "",
      dailyStartTime: "",
      dailyEndTime: "",
      duration: "",
      isOverlapping: true,
      isWholeDay: false,
    });
    setIsQuickModalOpen(true);
  };

  const handleCloseQuickModal = () => {
    setIsQuickModalOpen(false);
    setQuickData({
      startDate: "",
      endDate: "",
      dailyStartTime: "",
      dailyEndTime: "",
      duration: "",
      isOverlapping: true,
      isWholeDay: false,
    });
  };

  const generateTimeSlots = (): TimeSlot[] => {
    const generated: TimeSlot[] = [];
    const start = new Date(quickData.startDate);
    const end = new Date(quickData.endDate);

    if (quickData.isWholeDay) {
      for (let date = new Date(start); date <= end; date.setDate(date.getDate() + 1)) {
        const dateStr = date.toISOString().split("T")[0];
        generated.push({
          id: crypto.randomUUID(),
          date: dateStr,
          startTime: "00:00",
          endTime: "23:59",
        });
      }
      return generated;
    }

    const durationHours = parseFloat(quickData.duration);
    const intervalHours = quickData.isOverlapping ? durationHours / 2 : durationHours;

    const [dailyStartHour, dailyStartMin] = quickData.dailyStartTime.split(":").map(Number);
    const [dailyEndHour, dailyEndMin] = quickData.dailyEndTime.split(":").map(Number);
    const dailyStartMinutes = (dailyStartHour * 60) + dailyStartMin;
    const dailyEndMinutes = (dailyEndHour * 60) + dailyEndMin;

    for (let date = new Date(start); date <= end; date.setDate(date.getDate() + 1)) {
      const dateStr = date.toISOString().split("T")[0];

      let currentMinutes = dailyStartMinutes;
      while (currentMinutes + (durationHours * 60) <= dailyEndMinutes) {
        const startHour = Math.floor(currentMinutes / 60);
        const startMin = currentMinutes % 60;
        const endMinutes = currentMinutes + (durationHours * 60);
        const endHour = Math.floor(endMinutes / 60);
        const endMin = endMinutes % 60;

        const startTime = `${String(startHour).padStart(2, "0")}:${String(startMin).padStart(2, "0")}`;
        const endTime = `${String(endHour).padStart(2, "0")}:${String(endMin).padStart(2, "0")}`;

        generated.push({
          id: crypto.randomUUID(),
          date: dateStr,
          startTime,
          endTime,
        });

        currentMinutes += intervalHours * 60;
      }
    }

    return generated;
  };

  const handleGenerateQuickSlots = (onGenerate: (slots: TimeSlot[]) => void) => {
    const slots = generateTimeSlots();
    onGenerate(slots);
    handleCloseQuickModal();
  };

  const isQuickFormValid = quickData.isWholeDay
    ? quickData.startDate && quickData.endDate
    : quickData.startDate &&
    quickData.endDate &&
    quickData.dailyStartTime &&
    quickData.dailyEndTime &&
    quickData.duration;

  return {
    isQuickModalOpen,
    quickData,
    setQuickData,
    handleOpenQuickModal,
    handleCloseQuickModal,
    handleGenerateQuickSlots,
    isQuickFormValid,
  };
};

export const QuickSlotGenerator = ({
  startDate,
  endDate,
  dailyStartTime,
  dailyEndTime,
  duration,
  isOverlapping,
  isWholeDay,
  onStartDateChange,
  onEndDateChange,
  onDailyStartTimeChange,
  onDailyEndTimeChange,
  onDurationChange,
  onOverlappingChange,
  onWholeDayChange,
}: QuickSlotGeneratorProps) => {
  return (
    <>
      <div className="mt-3">
        <label htmlFor="start-date">Start date</label>
        <input
          id="start-date"
          type="date"
          className="form-control"
          value={startDate}
          onChange={(e) => {
            onStartDateChange(e.target.value);
          }}
        />
      </div>

      <div className="mt-3">
        <label htmlFor="end-date">End date</label>
        <input
          id="end-date"
          type="date"
          className="form-control"
          value={endDate}
          onChange={(e) => {
            onEndDateChange(e.target.value);
          }}
        />
      </div>

      <div className="mt-3">
        <div className="form-check">
          <input
            id="whole-day"
            type="checkbox"
            className="form-check-input"
            checked={isWholeDay}
            onChange={(e) => {
              const checked = e.target.checked;
              onWholeDayChange(checked);
              if (checked) {
                onDailyStartTimeChange("00:00");
                onDailyEndTimeChange("23:59");
              } else {
                onDailyStartTimeChange("");
                onDailyEndTimeChange("");
              }
            }}
          />
          <label
            className="form-check-label"
            htmlFor="whole-day">
            Whole day
          </label>
        </div>
        <small className="text-muted">
          {isWholeDay
            ? "Slots will cover the entire day (00:00 - 23:59)"
            : "Specify custom time range for daily slots"}
        </small>
      </div>

      {!isWholeDay ? <>
        <div className="mt-3">
          <label htmlFor="daily-start-time">Daily start time</label>
          <input
            id="daily-start-time"
            type="time"
            className="form-control"
            value={dailyStartTime}
            onChange={(e) => {
              onDailyStartTimeChange(e.target.value);
            }}
          />
        </div>

        <div className="mt-3">
          <label htmlFor="daily-end-time">Daily end time</label>
          <input
            id="daily-end-time"
            type="time"
            className="form-control"
            value={dailyEndTime}
            onChange={(e) => {
              onDailyEndTimeChange(e.target.value);
            }}
          />
        </div>
      </> : null}

      {!isWholeDay ? <>
        <div className="mt-3">
          <label htmlFor="duration">Slot duration (hours)</label>
          <select
            id="duration"
            className="form-control"
            value={duration}
            onChange={(e) => {
              onDurationChange(e.target.value);
            }}>
            <option value="">Select duration</option>
            <option value="0.5">30 minutes</option>
            <option value="1">1 hour</option>
            <option value="1.5">1.5 hours</option>
            <option value="2">2 hours</option>
            <option value="2.5">2.5 hours</option>
            <option value="3">3 hours</option>
            <option value="4">4 hours</option>
          </select>
        </div>

        <div className="mt-3">
          <div className="form-check">
            <input
              id="overlapping"
              type="checkbox"
              className="form-check-input"
              checked={isOverlapping}
              onChange={(e) => {
                onOverlappingChange(e.target.checked);
              }}
            />
            <label
              className="form-check-label"
              htmlFor="overlapping">
              Allow overlapping time slots
            </label>
          </div>
          <small className="text-muted">
            {isOverlapping
              ? "Slots will overlap (e.g., 16-18, 17-19, 18-20) for maximum flexibility"
              : "Slots will be consecutive (e.g., 16-18, 18-20, 20-22) with no overlap"}
          </small>
        </div>
      </> : null}

      <div className="alert alert-info mt-3">
        <small>
          This will generate time slots for each day in the date range within the
          specified time window.
        </small>
      </div>
    </>
  );
};
