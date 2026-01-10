export interface Calendar {
  id: string;
  title: string;
  description?: string;
  location?: string;
  acceptResponsesUntil?: string;
  createdAt: string;
  updatedAt: string;
}

export interface TimeSlot {
  id: string;
  slotDate: string;
  startTime: string;
  endTime: string;
  createdAt: string;
  updatedAt: string;
}
