import dayjs from "dayjs";
import { emptyApi } from "./emptyApi";

export const calendarsApi = emptyApi.injectEndpoints({
  endpoints: (builder) => ({
    createCalendar: builder.mutation<{
      id: string;
    }, {
      title: string;
      description?: string;
      location?: string;
      accept_responses_until?: string;
      password?: string;
    }>({
      query: (body) => ({
        url: "/calendars",
        method: "POST",
        body: {
          title: body.title,
          description: body.description || undefined,
          location: body.location || undefined,
          password: body.password || undefined,
          accept_responses_until: body.accept_responses_until
            ? dayjs(body.accept_responses_until).toISOString()
            : undefined,
        }
      }),
    }),
    createCalendarTimeSlots: builder.mutation<undefined, {
      calendar_id: string;
      time_slots: {
        start_date: string,
        end_date: string
      }[]
    }>({
      query: (body) => ({
        url: `/calendars/${body.calendar_id}/time-slots`,
        method: "POST",
        body: {
          time_slots: body.time_slots.map(slot => {
            const parseLocalDateTime = (dateTimeStr: string) => {
              const [datePart, timePart] = dateTimeStr.split("T");
              const [year, month, day] = datePart.split("-").map(Number);
              const [hour, minute, second] = timePart.split(":").map(Number);
              return new Date(year, month - 1, day, hour, minute, second || 0);
            };

            const startLocal = parseLocalDateTime(slot.start_date);
            const endLocal = parseLocalDateTime(slot.end_date);

            return {
              start_date: startLocal.toISOString(),
              end_date: endLocal.toISOString(),
            };
          }),
        }
      }),
    }),
  })
});

export const {
  useCreateCalendarMutation,
  useCreateCalendarTimeSlotsMutation
} = calendarsApi;