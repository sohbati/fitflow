package com.caa.util;

import ir.huri.jcal.JalaliCalendar;
import java.util.Calendar;
import java.util.Date;
import java.util.GregorianCalendar;

/**
 * Created by Reza on 16/08/2018.
 */
public class DateUtil {

    public static String  getShamsiDate(Date d) {
        Calendar calendar = Calendar.getInstance();
        calendar.setTime(d);
        JalaliCalendar j = new JalaliCalendar(new GregorianCalendar(
                calendar.get(Calendar.YEAR), calendar.get(Calendar.MONTH), calendar.get(Calendar.DAY_OF_MONTH)));
        return j.getYear() + "/" + j.getMonth() + "/" + j.getDay();

//        PersianCalendar persianCalendar1 = new PersianCalendar(d);
//        String s = persianCalendar1.get(PersianCalendar.YEAR) + "/" +
//                persianCalendar1.get(PersianCalendar.MONTH) + "/" +
//                persianCalendar1.get(PersianCalendar.DAY_OF_MONTH) ;
//        return s;
    }

    public static Date getGregorianDate(String d) {
        if (d == null || d.trim().length() == 0) {
            return new Date(0);
        }
        String arr[] = d.trim().split("/");
        JalaliCalendar jalaliCalendar =
                new JalaliCalendar(Integer.parseInt(arr[0]), Integer.parseInt(arr[1]), Integer.parseInt(arr[2]));
        return jalaliCalendar.toGregorian().getTime();
    }
}
